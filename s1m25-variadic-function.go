package main

import "fmt"

func main() {
	fmt.Println("total", sumAll(1, 2, 3, 4, 5))

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("total", sumAll(numbers...))
}

func sumAll(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}
