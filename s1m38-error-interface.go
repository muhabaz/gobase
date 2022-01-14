package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(Pembagian(5, 4))
	fmt.Println(Pembagian(5, 0))
}

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagian salah")
	}
	return nilai / pembagi, nil
}
