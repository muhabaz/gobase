package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

type HasName interface {
	GetName() string
}

func main() {
	customer := Customer{
		Name:    "John",
		Address: "tangerang",
		Age:     20,
	}
	SayHello(customer)
}

func (c Customer) GetName() string {
	return c.Name
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}
