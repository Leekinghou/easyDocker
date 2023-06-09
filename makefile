CMD=go
BIN_PATH=bin
SRC_PATH=src

# all: 一个指令调用多个命令
all: clean build install

build: 
	$(CMD) build -o $(BIN_PATH)/docker $(SRC_PATH)/*

install:
	cp bin/docker /usr/bin/docker
	cp bin/docker /usr/local/bin/docker

uninstall:
	rm -rf /usr/bin/docker /usr/local/bin/docker
	rm -rf bin/docker

clean: uninstall