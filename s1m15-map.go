package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "John",
		"address": "Tangerang",
	}
	person["email"] = "mail@mail.mail"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	delete(person, "email")
	fmt.Println(person)

	book := make(map[string]string)
	fmt.Println(book)
}
