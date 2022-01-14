package main

import "fmt"

func main() {
	fmt.Println("nice")
	run(5)
	run(0)
}

func log() {
	fmt.Println("done")

	message := recover()
	if message != nil {
		fmt.Println("error:", message)
	}
}

func run(value int) {
	defer log()
	fmt.Println("running")
	if value == 0 {
		panic("illegal parameter")
	}
	result := 10 / value
	fmt.Println("result", result)
}
