package main

import "fmt"

type Test struct {
	Name string
	Age  int
}

func (test *Test) GetName() string {
	return test.Name
}

func (test *Test) GetAge() int {
	return test.Age
}

type Test1 struct {
	Test
}

func main() {
	test := Test{
		Name: "test",
		Age:  18,
	}
	fmt.Printf("Name: %s\nAge: %d\n", test.GetName(), test.GetAge())
}
