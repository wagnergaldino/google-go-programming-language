package main

import "fmt"

func foo() {
	defer bar()
	fmt.Println("printing from foo")
}

func bar() {
	fmt.Println("printing from bar")
}

func main() {
	foo()
}
