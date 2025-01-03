package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	s2 := make([]int, 4)
	copy(s2, s)
	fmt.Println(s2)
}
