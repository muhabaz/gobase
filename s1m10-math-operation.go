package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := a * b

	fmt.Println(c)

	c += b
	fmt.Println(c)
	fmt.Println(-c)

}
