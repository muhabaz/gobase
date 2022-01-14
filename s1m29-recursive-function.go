package main

import "fmt"

func main() {
	fmt.Println("facto", factorialLoop(5))
	fmt.Println("facto v2", factorial(5))
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorial(value int) int {
	if value <= 1 {
		return 1
	} else {
		return value * factorial(value-1)
	}
}
