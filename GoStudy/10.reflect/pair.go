package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u *User) Call() {
	fmt.Printf("I'm %s, I'm %d years old\n", u.Name, u.Age)
}

func Do(input interface{}) {

	inputtype := reflect.TypeOf(input)
	fmt.Printf("input type is %v\n", inputtype.Name())

	inputvalue := reflect.ValueOf(input)
	fmt.Printf("input value is %v\n", inputvalue)

	for i := 0; i < inputtype.NumMethod(); i++ {
		method := inputtype.Method(i)
		fmt.Printf("method name is %v\n", method.Name)
		fmt.Printf("method type is %v\n", method.Type)
		fmt.Printf("method value is %v\n", method.Func)
		fmt.Printf("method value is %v\n", method.Index)
		fmt.Printf("method value is %v\n", method.PkgPath)
	}
}

func main() {
	var i interface{} = "Hello"
	if _, ok := i.(string); ok {
		fmt.Println("It's a string:")
	} else {
		fmt.Println("It's not a string")
	}
	user := &User{Name: "John", Age: 25}
	Do(user)
}
