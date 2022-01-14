package main

import "fmt"

type Man struct {
	Name string
}

func main() {
	abdul := Man{"abdul"}
	abdul.Married()

	fmt.Println(abdul)
}

func (man *Man) Married() {
	man.Name = "Mr " + man.Name
}
