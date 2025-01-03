package main

import (
	"fmt"
)

var (
	gA = 100
	gB = 100
)

func main() {
	var a = 100
	fmt.Printf("type of a = %T\n", a)

	e := 100
	fmt.Printf("value of e = %v\n", gA)
	fmt.Printf("type of e = %T\n", e)
}
