# ACG - System Tooling with Go

This repository contains some simple scripts about what I learned from A Cloud Guru - System Tooling with Go.

## Environment

* Download Go and set up environment variable based on your need

	```sh
	mkdir -p $HOME/go/{src,bin}
	export $GOPATH=$HOME/go
	export $PATH=$PATH:$GOPATH/bin
	```
* Install 'faith/go' pluggin on Vim editor and install 
	
	```
	vim +PluginInstall +qa
	vim +GoInstallBinaries +qa
	```

## Running Go

* Create a simple go app: ```vim main.go```

	```go
	package main
	
	import "fmt"
	
	func main() {
	  fmt.Println("vim-go")
	}
	```

* Run the go app:

	```go
	go run main.go
	```
* Build a go executable

	```go
	go build main.go
	./main
	```
