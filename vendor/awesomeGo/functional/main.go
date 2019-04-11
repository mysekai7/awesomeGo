package main

import (
	"awesomeProject1/functional/fib"
	"bufio"
	"fmt"
	"io"
	"strings"
)

type intGen func() int

////1,1,2,3,5,8,13,...
////  a,b
////	  a,b
////斐波那契数列
//func Fibonacci() intGen {
//	a, b := 0, 1
//	return func() int {
//		a, b = b, a+b
//		return a
//	}
//}

//函数式接口
//函数类型实现Read接口, 函数类型作为接受者
//Read(p []byte) (n int, err error)
func (g intGen) Read(p []byte) (n int, err error) {
	next := g() //下一个元素
	if next > 100 {
		return 0, io.EOF //读到头
	}
	s := fmt.Sprintf("%d\n", next) //转成字符串

	//TODO: p存储太小的问题
	return strings.NewReader(s).Read(p) //代理read
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {

	f := fib.Fibonacci()
	//fmt.Println(f()) //1
	//fmt.Println(f()) //1
	//fmt.Println(f()) //2
	//fmt.Println(f()) //3
	//fmt.Println(f()) //5
	//fmt.Println(f()) //8

	printFileContents(f)

}
