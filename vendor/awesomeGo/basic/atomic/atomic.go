package main

import (
	"fmt"
	"sync"
	"time"
)

//线程安全
type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increament() {
	fmt.Println("safe increment")

	//代码块保护,使用匿名函数
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increament()
	go func() {
		a.increament()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
