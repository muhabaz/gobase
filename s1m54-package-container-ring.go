package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// https://pkg.go.dev/container/ring
	data := ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	fmt.Println(data)
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
