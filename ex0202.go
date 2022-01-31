package main

import "fmt"

func main() {
	a := 2
	b := 5

	g := a == b
	h := a <= b
	i := a >= b
	j := a != b
	k := a < b
	l := a > b

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}
