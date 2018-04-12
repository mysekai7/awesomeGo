package main

import (
	"awesomeProject1/retriever/mock"
	"awesomeProject1/retriever/real"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"hello word"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
