package main

import (
	"fmt"
)

func foo(a int, b int) (r1, r2 int) {
	return a + b, a - b
}
func main() {
	ret1, ret2 := foo(1, 2)
	fmt.Println("foo", ret1, ret2)

}
