package main

import "fmt"

func main() {
	c, s := getMultiple(5, 3)
	fmt.Println(c, s)

	d, _ := getMultiple(5, 7)
	fmt.Println(d)
}

func getMultiple(a int, b int) (int, string) {
	c := a * b
	return c, "ok"
}
