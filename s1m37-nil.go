package main

import "fmt"

func main() {
	var person map[string]string = nil
	person = NewMap("abdul")
	fmt.Println(person)
}

func NewMap(name string) map[string]string {
	if name != "" {
		return map[string]string{"name": name}
	}
	return nil
}
