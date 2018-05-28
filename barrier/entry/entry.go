package main

import (
	"awesomeGo/barrier"
	"fmt"
)

func main() {

	waitNum := 4

	done := make(chan bool, waitNum)
	br := barrier.NewBarrier(waitNum)

	for i := 0; i < waitNum; i++ {
		go worker(i, br, done)
	}

	for i := 0; i < waitNum; i++ {
		<-done
	}
}

func worker(i int, br *barrier.Barrier, done chan<- bool) {

	fmt.Printf("i:%d ---  A \n", i)
	br.Wait()
	fmt.Printf("i:%d ---  B \n", i)
	br.Wait()
	fmt.Printf("i:%d ---  C \n", i)

	done <- true
}
