package main

import "fmt"

func main() {
	i := 1965
	for {
		fmt.Println(i)
		if i == 2021 {
			break
		}
		i++
	}
}
