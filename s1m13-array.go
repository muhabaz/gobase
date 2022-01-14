package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Abdul"
	names[2] = "Azis"

	fmt.Println(names)
	fmt.Println(names[1])

	names[1] = "Abdula"
	fmt.Println(names[1])

	fmt.Println(len(names)) // panjang array, bukan data

	var values = [3]int{
		1, 2, 3,
	}

	fmt.Println(values)
	fmt.Println(values[1])

}
