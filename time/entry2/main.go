package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now().Add(time.Second * 3600 * 24 * 4)
	end := time.Now()
	fmt.Println(now.Sub(end).Hours())
}
