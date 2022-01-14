package main

import (
	"fmt"
	"math"
)

func main() {
	// https://pkg.go.dev/math
	fmt.Println("math", math.Round(1.7))
	fmt.Println("math", math.Round(1.3))
	fmt.Println("math", math.Floor(1.7))
	fmt.Println("math", math.Ceil(1.3))
}
