package main

import (
	"fmt"
)

func main() {
	cr := make(chan int)

	go func() {
		cr <- 55
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
