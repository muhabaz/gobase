package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"20"`
}

func main() {
	// https://pkg.go.dev/reflect
	sample := Sample{"abdul"}

	sampleTypes := reflect.TypeOf(sample)

	fmt.Println(sample)
	fmt.Println(sampleTypes.NumField())
	fmt.Println(sampleTypes.Field(0).Name)
	fmt.Println(sampleTypes.Field(0).Tag.Get("required"))
	fmt.Println(sampleTypes.Field(0).Tag.Get("max"))

	fmt.Println(IsValid(sample))
	fmt.Println(IsValid(Sample{}))
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}
