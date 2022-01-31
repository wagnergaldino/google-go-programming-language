package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 55
	}()

	fmt.Println(<-c)
}
