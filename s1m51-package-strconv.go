package main

import (
	"fmt"
	"strconv"
)

func main() {
	// https://pkg.go.dev/strconv
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("result", boolean)
	} else {
		fmt.Println("error", err)
	}

	number, err := strconv.ParseInt("30000", 10, 64)
	if err == nil {
		fmt.Println("result", number)
	} else {
		fmt.Println("error", err)
	}
}
