package main

import "fmt"

func main() {
	result := random()

	// resultStr := result.(string)
	// fmt.Println(resultStr)

	// resultBool := result.(bool)
	// fmt.Println(resultBool)

	switch value := result.(type) {
	case string:
		fmt.Println(value)
	case bool:
		fmt.Println(value)
	default:
		fmt.Println(value)
	}
}

func random() interface{} {
	// return "random"
	return true
}
