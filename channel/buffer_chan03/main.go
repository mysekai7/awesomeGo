package main

import (
	"fmt"
	"log"
	"time"
)

func read(ch chan int) {
	var i int
	for {
		select {
		case i = <-ch:
			time.Sleep(time.Second * 2)
			fmt.Println("successfully, read", i, "from ch")
		}

	}
	//close(ch)
}

func main() {
	ch := make(chan int, 100)
	go read(ch)
	i := 0
	for i < 1000 {
		ch <- i
		log.Println("ch <- ", i)
		i++
	}
}
