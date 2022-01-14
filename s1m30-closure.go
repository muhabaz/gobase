package main

import "fmt"

func main() {
	counter := 0
	increment := func() {
		counter := 0
		fmt.Println("inc", counter)
		counter++
		fmt.Println("inc", counter)
	}

	increment()
	fmt.Println("waw", counter)
}
