package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	//for {
	//	n, ok := <-c	//判断chan是否已close
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("worker %d received %d \n",
	//		id, n)
	//}
	for n := range c { //判断chan是否已close
		fmt.Printf("worker %d received %d \n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3) //缓冲区，提升性能
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3) //缓冲区，提升性能
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
	//bufferedChannel()
	//channelClose()
}
