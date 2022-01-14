package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	customer := Customer{Name: "John", Address: "tangerang", Age: 20}
	customer.sayHello()

}

func (customer Customer) sayHello() {
	fmt.Println("Hello", customer.Name)
}
