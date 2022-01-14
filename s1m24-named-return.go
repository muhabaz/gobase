package main

import "fmt"

func main() {
	firstName, middleName, lastName := getName()
	fmt.Println(firstName, middleName, lastName)
}

func getName() (firstName, middleName, lastName string) {
	firstName = "muhammad"
	middleName = "abdul"
	lastName = "azis"
	return
}
