package main

import (
	"fmt"
)

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	// ch := make(chan int)

	// go func(ch chan int) {
	// 	fmt.Println("data:", <-ch)
	// }(ch)
	// ch <- 5
	// time.Sleep(1 * time.Microsecond)

	sendch := make(chan int)
	go sendData(sendch)
	fmt.Println(<-sendch)
}
