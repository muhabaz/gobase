package main

import "fmt"

func main() {
	goodBye := getGoodBye

	fmt.Println(goodBye("abdul"))
}

func getGoodBye(name string) string {
	return "Good bye " + name
}
