package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) { //race condition
			for {
				//fmt.Printf("hello from "+
				//	"goroutine %d\n", i)
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Microsecond)
	fmt.Println(a)
}
