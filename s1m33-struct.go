package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	customer := Customer{Name: "John", Address: "tangerang", Age: 20}
	fmt.Println(customer)
	fmt.Println(customer.Address)

}
