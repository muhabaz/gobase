package main

import (
	"container/list"
	"fmt"
)

func main() {
	// https://pkg.go.dev/container/list
	data := list.New()
	data.PushBack("muhammad")
	data.PushBack("abdul")
	data.PushBack("azis")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Front().Next().Value)
	fmt.Println(data.Back().Value)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
