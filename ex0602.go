package main

import "fmt"

func foo(x ...int) int {
	valor := 0
	for _, v := range x {
		valor += v
	}
	return valor
}

func bar(x []int) int {
	valor := 0
	for _, v := range x {
		valor += v
	}
	return valor
}

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(foo(a...))
	fmt.Println(bar(a))
}
