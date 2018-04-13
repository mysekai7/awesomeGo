package main

import (
	"awesomeProject1/retriever/mock"
	"awesomeProject1/retriever/real"
	"fmt"
	"time"
)

//接口定义
type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever

	//值类型，即可使用值类型也可以使用指针类型
	r = mock.Retriever{"hello word"}
	inspect(r)

	/*
		retriever/main.go:24:4: cannot use real.Retriever literal (type real.Retriever) as type Retriever in assignment:
		real.Retriever does not implement Retriever (Get method has pointer receiver)
	*/
	//指针类型, 指针接受者必须指针类型
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	inspect(r)

	//Type assertion， 查看指针变量
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	//realRetriever := r.(*real.Retriever) // .(type)获取具体类型
	//fmt.Println(realRetriever.Timeout)

	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}

	fmt.Printf("%T %v\n", r, r)
}
