package main

import (
	"fmt"
)

func hello(done chan int, number int) {
	fmt.Println("hello world goroutine")
	done <- number
}

func main() {
	done := make(chan int)

	for i := 0; i < 10; i++ {
		go hello(done, i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("i: ", <-done)
	}

	fmt.Println("main function")
}
