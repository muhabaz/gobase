package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Counter", counter)
		counter++
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Start", i)
	}

	var months = []string{
		"Jan", "Feb", "Mar",
		"Apr", "May", "Jun",
		"Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec",
	}
	fmt.Println(months)
	for i := 0; i < len(months); i++ {
		fmt.Println(months[i])
	}

	// range
	for i, month := range months {
		fmt.Println(i, month)
	}
	for _, month := range months {
		fmt.Println(month)
	}

	person := map[string]string{
		"name":    "John",
		"address": "Tangerang",
	}
	for key, value := range person {
		fmt.Println(key, value)
	}
}
