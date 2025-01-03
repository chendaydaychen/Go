package main

import (
	"fmt"
	"time"
)

func main() {
	/*c := make(chan int)

	go func() {
		defer fmt.Println("defer in goroutine")
		fmt.Println("goroutine")

		c <- 666
	}()

	num := <-c
	fmt.Println("main", num) */
	c := make(chan int, 3)
	go func() {

		defer fmt.Println("defer in goroutine")
		for i := 0; i < 6; i++ {
			c <- i
			fmt.Println(len(c), cap(c))
			fmt.Println("goroutine", i)
		}
	}()

	time.Sleep(time.Second)

	for i := 0; i < 6; i++ {
		fmt.Println("num", <-c)
	}
	fmt.Println("main")
}
