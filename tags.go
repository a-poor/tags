package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID          int     `csv:"user_id"`
	FirstName   string  `csv:"first_name"`
	LastName    string  `csv:"last_name,omitempty"`
	CoolnessPct float32 `csv:"coolness_pct"`
	IsCool      bool    `csv:"is_cool"`
}

func main() {
	u := User{1, "Tim", "Tomson", 0.9, true}
	fmt.Printf("%+v\n", u)

	ut := reflect.TypeOf(u)
	fmt.Printf("%+v\n", ut)
	for i := 0; i < ut.NumField(); i++ {
		field := ut.Field(i)
		fmt.Printf("[%d] %15q => %q\n", i, field.Name, field.Tag)
		csv, ok := field.Tag.Lookup("csv")
		fmt.Printf("csv set: %t   csv value: %q\n", ok, csv)
		fmt.Println()
	}
}
