package main

import (
	"fmt"
)

func main() {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("todos os numeros", i)

	s := soma(i...)
	fmt.Println("soma de todos os numeros", s)

	s2 := pares(soma, i...)
	fmt.Println("soma dos numeros pares", s2)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func pares(f func(i ...int) int, j ...int) int {
	var x []int
	for _, v := range j {
		if v%2 == 0 {
			x = append(x, v)
		}
	}
	return f(x...)
}
