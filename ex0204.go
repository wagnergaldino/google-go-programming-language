package main

import "fmt"

var x int = 25

func main() {
	fmt.Printf("%d - %b - %x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d - %b - %x\n", y, y, y)
}
