package main

import (
	"fmt"
	"os"
)

func main() {
	// https://pkg.go.dev/os
	args := os.Args
	hostname, err := os.Hostname()
	gopath := os.Getenv("GOPATH")
	fmt.Println(args)
	fmt.Println(hostname)
	fmt.Println(err)
	fmt.Println(gopath)
}
