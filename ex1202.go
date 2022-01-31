package main

import "fmt"

func main() {
	i, j := 6, 7

	f := func(x int) (y int, isEven bool) {
		y = x / 2
		if x%2 == 0 {
			isEven = true
		} else {
			isEven = false
		}
		return
	}

	a, b := f(i)
	fmt.Println(i, "dividido por 2 =", a, "-", i, "é par?", b)

	c, d := f(j)
	fmt.Println(j, "dividido por 2 =", c, "-", j, "é par?", d)
}
