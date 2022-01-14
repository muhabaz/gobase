package main

import "fmt"

func main() {
	name := "azis"
	if name == "abdul" {
		fmt.Println("hello abdul")
	} else if name == "azis" {
		fmt.Println("hello azis")
	} else {
		fmt.Println("who are you?")
	}

	// short statement
	if length := len(name); length < 10 {
		fmt.Println("ok jos")
	}
}
