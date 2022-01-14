package main

import "fmt"

func main() {
	name1 := "abdul"
	name2 := "abdul"
	name3 := "azis"

	result := name1 == name2
	result2 := name1 == name3
	result3 := name1 < name3

	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
}
