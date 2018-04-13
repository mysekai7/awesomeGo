package main

import "fmt"

//闭包
func adder() func(int) int {
	sum := 0 //自由变量
	return func(v int) int {
		sum += v //对sum引用
		return sum
	}
}

func main() {

	a := adder()
	for i := 0; i < 10; i++ {
		//fmt.Println(a(i))
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	}
}
