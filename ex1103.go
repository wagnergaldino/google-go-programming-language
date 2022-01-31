package main

import (
	"fmt"
)

type customErr struct {
	s string
}

func (ce customErr) Error() string {
	return ce.s
}

func main() {
	c1 := customErr{}
	c1.s = "customErr"
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
