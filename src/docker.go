package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	fmt.Printf("Process => %v [%d]\n", os.Args, os.Getpid())
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("have not defined")
	}
}

func run() {
	cmd := exec.Command(os.Args[2], append([]string{"child"}, os.Args[2])...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS, // linux特有
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func child() {
	cmd := exec.Command(os.Args[2])
	syscall.Sethostname([]byte("container"))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
