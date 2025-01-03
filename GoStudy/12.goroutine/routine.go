package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		fmt.Println("i=", i)
		i++
		time.Sleep(time.Second)
	}

}

func main() {
	go newTask()

	i := 0
	for {
		fmt.Println("i=", i)
		i++
		time.Sleep(time.Second)
	}
}
