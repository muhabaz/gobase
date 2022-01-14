package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address := Address{
		City:     "a",
		Province: "s",
		Country:  "d",
	}
	ChangeAddress(&address)

	fmt.Println(address)
}

func ChangeAddress(address *Address) {
	address.Country = "japan"
}
