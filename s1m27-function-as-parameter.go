package main

import (
	"fmt"
	"strings"
)

type Filter func(string) string

func main() {
	sayHelloWithFilter("abdul", upperFilter)
	sayHelloWithFilter("anjing", spamFilter)
}

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	}
	return name
}

func upperFilter(name string) string {
	return strings.ToUpper(name)
}
