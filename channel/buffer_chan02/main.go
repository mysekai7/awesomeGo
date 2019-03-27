package main

import (
	// "fmt"
	"log"
	"time"
)

func write(ch chan int) {
	// for i := 0; i < 5; i++ {
	// 	ch <- i
	// 	fmt.Println("successfully wrote", i, "to ch")
	// }
	// close(ch)

	for {
		select {
		case c := <- ch:
			log.Println("ch: ", c)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ch := make(chan int, 100)
	go write(ch)
	//time.Sleep(2 * time.Second)
	// for v := range ch {
	// 	fmt.Println("read value", v, "from ch")
	// 	time.Sleep(2 * time.Second)
	// }
	for i:=0; i<200; i++ {
		ch <- i
		log.Println("i: ", i)
	}
	time.Sleep(100 * time.Second)
}
