package main

import (
	"fmt"
	"time"
)

func process(i int, ch chan bool) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	ch <- true
}

func main() {
	no := 3
	done := make(chan bool)
	for i := 0; i < no; i++ {
		go process(i, done)
	}
	for i := 0; i < no; i++ {
		<-done
	}
	fmt.Println("All go routines finished executing")
}
