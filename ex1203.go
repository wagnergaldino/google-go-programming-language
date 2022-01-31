package main

import "fmt"

func main() {
	x := biggest(14, 2, 5, -2, 7, 11)
	fmt.Println("The biggest number is", x)
}

func biggest(x ...int) int {
	var y int
	for i, v := range x {
		if i == 0 {
			y = v
		} else {
			if v > y {
				y = v
			}
		}
	}
	return y
}
