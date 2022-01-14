package main

import (
	"fmt"
	"regexp"
)

func main() {
	// https://pkg.go.dev/regexp
	// https://github.com/google/re2/wiki/Syntax
	regex := regexp.MustCompile("a([a-z])s")

	fmt.Println(regex.MatchString("aus"))
	fmt.Println(regex.MatchString("ais"))
	fmt.Println(regex.MatchString("abi"))
}
