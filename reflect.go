package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Akkshay", Age: 20}

	// type & value of struct
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	fmt.Println("Type:", t)
	fmt.Println("Value:", v)

	// iterate over fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("Field: %s, Type: %s, Value: %v\n", field.Name, field.Type, value)
	}

	// val of feild with pointer
	pPtr := reflect.ValueOf(&p).Elem()
	pPtr.FieldByName("Age").SetInt(21)
	fmt.Println("Updated struct:", p)
}
