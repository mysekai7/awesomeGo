package main

import (
	"awesomeProject1/queue"

	"fmt"

	"golang.org/x/tools/container/intsets"
)

func main() {

	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	s := intsets.Sparse{}
	s.Insert(1)

	fmt.Println(s)
}
