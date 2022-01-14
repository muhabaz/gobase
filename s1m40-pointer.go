package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		City:     "tangerang",
		Province: "banten",
		Country:  "indonesia",
	}
	address2 := &address1
	address3 := &address1
	address4 := new(Address)

	address2.City = "serang"

	*address2 = Address{
		City:     "jogja",
		Province: "jogja",
		Country:  "indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4.City = "jakarta"
	fmt.Println(address4)
}
