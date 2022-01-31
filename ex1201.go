package main

import "fmt"

func main() {
	i, j := 4, 5
	a, b := myFunc(i)
	fmt.Println(i, "dividido por 2 =", a, "-", i, "é par?", b)

	c, d := myFunc(j)
	fmt.Println(j, "dividido por 2 =", c, "-", j, "é par?", d)
}

func myFunc(x int) (y int, isEven bool) {
	y = x / 2
	if x%2 == 0 {
		isEven = true
	} else {
		isEven = false
	}
	return
}
