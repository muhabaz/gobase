package main

import "fmt"

func main() {
	type NoKTP string

	var ktpAzis NoKTP = "1234567890"

	fmt.Println(ktpAzis)
}
