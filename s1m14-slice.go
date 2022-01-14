package main

import "fmt"

func main() {
	// pointer, length, capacity
	var months = [...]string{
		"Jan", "Feb", "Mar",
		"Apr", "May", "Jun",
		"Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "ubah"
	// fmt.Println(slice1)

	// slice1[0] = "mei"
	// fmt.Println(slice1)
	// fmt.Println(months)

	// days := [...]string{
	// 	"sunday", "monday", "tuesday",
	// 	"wednesday", "thursday", "friday",
	// 	"saturday",}

	var slice2 = months[10:]
	// var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "abdul")
	fmt.Println(slice3)
	slice3[1] = "desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "abdul"
	newSlice[1] = "azis"
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// array [...], slice []
}
