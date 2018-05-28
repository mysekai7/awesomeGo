package main

import (
	"reflect"
	"fmt"
	"time"
	"runtime"
)

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func Decorator(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			t := time.Now()

			//fmt.Println("before")
			out = targetFunc.Call(in)
			//fmt.Println("after")

			fmt.Printf("--- Time Elapsed (%s): %v ---\n",
				getFunctionName(fn), time.Since(t))

			return
		})

	decoratedFunc.Set(v)
	return
}

func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b, c)
	return a + b + c
}

func bar(a, b string) string {
	fmt.Printf("%s, %s \n", a, b)
	return a + b
}

func Sum1(start, end int64) int64 {
	var sum int64
	sum = 0
	if start > end {
		start, end = end, start
	}
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}

func Sum2(start, end int64) int64 {
	if start > end {
		start, end = end, start
	}
	return (end - start + 1) * (end + start) / 2
}

func main() {

	//type MyFoo func(int, int, int) int
	//var myfoo MyFoo
	//Decorator(&myfoo, foo)
	//myfoo(1, 2, 3)

	//myfoo := foo
	//Decorator(&myfoo, foo)
	//myfoo(1,2,3)

	//mybar := bar
	//Decorator(&mybar, bar)
	//mybar("hello", "world!")


	mySum1 := Sum1
	Decorator(&mySum1, Sum1)

	mySum2 := Sum2
	Decorator(&mySum2, Sum2)

	fmt.Printf("%d, %d\n", mySum1(-10000, 10000000), mySum2(-10000, 10000000))
	
}
