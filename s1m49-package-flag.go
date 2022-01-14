package main

import (
	"flag"
	"fmt"
)

func main() {
	// https://pkg.go.dev/flag
	host := flag.String("host", "localhost", "host to connect to")
	username := flag.String("username", "root", "username to connect to")
	password := flag.String("password", "root", "password to connect to")

	flag.Parse()

	fmt.Println("Connecting to", *host, *username, *password)
}
