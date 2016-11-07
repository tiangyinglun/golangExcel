package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
}

type Admin struct {
	User
	title string
}

func main() {
	var u Admin
	t := reflect.TypeOf(u)
	fmt.Println(t)
	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f)
		fmt.Println(f.Name, f.Type)
	}
}
