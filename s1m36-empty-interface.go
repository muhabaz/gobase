package main

import "fmt"

func main() {
	data := Ups(2)
	fmt.Println(data)
}

func Ups(i int) interface{} {
	if i == 0 {
		return false
	} else if i == 1 {
		return true
	}
	return "Ups"
}
