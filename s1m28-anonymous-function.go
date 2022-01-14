package main

import "fmt"

type Blacklist func(string) bool

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("abdul", blacklist)
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("blocked, cannot use", name)
	} else {
		fmt.Println("welcome", name)
	}
}
