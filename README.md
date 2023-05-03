# version 0.1

# version 0.2
实现namespace  

## 什么是namespace
Namespace 是 Linux 内核用来**隔离内核资源**的方式，是对**全局系统资源**的一种封装隔离，使得处于不同 namespace 的进程拥有独立的全局系统资源，改变一个 namespace 中的系统资源只会影响当前 namespace 里的进程，对其他 namespace 中的进程没有影响，但实际上它们共享同一个物理主机的资源，这样，不同的容器就可以在同一台主机上运行不同的应用程序，而不会互相影响，从而实现更高效、更灵活的应用部署。

也就是说namespace可以隔离：进程ID、主机名、用户ID、文件名与网络访问相关的接口

Docker 实现了以下8种 namespace：

1. PID namespace：隔离进程树，使得每个容器内的进程 ID 看起来与主机上的进程 ID 不同，从而实现独立的进程空间。
2. NET namespace：隔离网络接口，使得每个容器内的网络接口看起来与主机上的网络接口不同，从而实现独立的网络空间。
3. MNT namespace：隔离文件系统，使得每个容器内的文件系统看起来与主机上的文件系统不同，从而实现独立的文件系统空间。
4. UTS namespace：隔离主机名，使得每个容器内的主机名看起来与主机上的主机名不同，从而实现独立的主机名空间。
5. IPC namespace：隔离进程间通信（IPC），使得每个容器内的 IPC 看起来与主机上的 IPC 不同，从而实现独立的 IPC 空间。
6. USER namespace：隔离用户和用户组，使得每个容器内的用户和用户组看起来与主机上的用户和用户组不同，从而实现独立的用户空间。
7. Control group(cgroup) namespace：隔离Cgroups根目录，限制了进程能使用的资源量 (CPU、 内存、 网络带宽等）
8. Time namespace：隔离系统时间

通过这些隔离机制，Docker 可以实现更加安全、高效、灵活的容器化部署。

## 实现原理：
```shell
sudo unshare --fork --pid --mount-proc bash
```
这条指令的作用是创建一个新的进程命名空间（PID namespace 和 Mount namespace），并在其中启动一个新的 Bash shell。在这个新的命名空间中，进程树、文件系统挂载、进程的PID等都是隔离的，与原来的命名空间是互相独立的。

具体来说，指令的各个参数的含义如下：  
`sudo`: 以超级用户权限执行指令。  
`unshare`: 创建一个新的命名空间，并将指定的进程与新的命名空间关联。  
`--fork`: 在新的命名空间中创建一个子进程。  
`--pid`: 创建一个新的PID命名空间，使得新的进程在这个命名空间中运行。  
`--mount-proc`: 在新的命名空间中，挂载一个新的/proc文件系统，使得新的进程可以看到这个文件系统中的进程信息。  
`bash`: 在新的命名空间中启动一个新的 Bash shell。  