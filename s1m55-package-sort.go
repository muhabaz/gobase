package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (u UserSlice) Len() int { return len(u) }

func (u UserSlice) Less(i, j int) bool {
	// return u[i].Name < u[j].Name
	return u[i].Age < u[j].Age
}

func (u UserSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	// https://pkg.go.dev/sort
	users := UserSlice{
		{"abdul", 20},
		{"azis", 23},
		{"budi", 14},
		{"adi", 43},
		{"ida", 22},
	}

	fmt.Println(users)
	sort.Sort(users)
	fmt.Println(users)

}
