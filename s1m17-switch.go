package main

import "fmt"

func main() {
	name := "abdul"

	switch name {
	case "abi":
		fmt.Println("hello abi")
	case "abdul":
		fmt.Println("hello abdul")
	default:
		fmt.Println("who are you?")
	}

	switch length := len(name); length <= 5 {
	case true:
		fmt.Println("mantap")
	case false:
		fmt.Println("waduh")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("mantap")
	case length > 5:
		fmt.Println("ok")
	case length <= 5:
		fmt.Println("no")
	}
}
