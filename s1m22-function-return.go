package main

import "fmt"

func main() {
	fmt.Println("total =", getTotal(5, 6))
}

func getTotal(a int, b int) int {
	return a + b
}
