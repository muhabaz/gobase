package main

import "fmt"

func main() {
	var nilai32 int32 = 327680
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai64)

	fmt.Println(nilai16)
	fmt.Println(nilai32)
	fmt.Println(nilai64)

	name := "Muhammad Abdul Azis"
	c := name[7]
	s := string(c)

	fmt.Println(s)
}
