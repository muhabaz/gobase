package main

import (
	"fmt"
	"time"
)

func main() {
	// https://pkg.go.dev/time
	now := time.Now()
	fmt.Println("now", now.Local())

	utc := time.Date(2021, 1, 12, 15, 30, 30, 10, time.UTC)
	fmt.Println("utc", utc)

	layout := "2006-01-02"
	// parse, _ := time.Parse(time.RFC3339, "2021-01-12")
	parse, _ := time.Parse(layout, "2021-01-12")
	fmt.Println("parse", parse)
}
