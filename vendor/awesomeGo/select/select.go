package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			//产生数据的随机时间
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c { //判断chan是否已close
		time.Sleep(time.Second * 1)
		fmt.Printf("worker %d received %d \n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {

	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second) //定时
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue: //一个值一个值给
			values = values[1:]
		case <-time.After(800 * time.Millisecond): //时间差,for循环相邻两个请求的时间差
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len= ", len(values))
		case <-tm: //定时结束
			fmt.Println("bye")
			return
		}
	}

}
