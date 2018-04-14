package main

import (
	"errors"
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occured:", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do : %v", r))
		}
	}()
	panic(errors.New("this is a error"))
}

func main() {
	tryRecover()
}
