package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	z := [][]string{x, y}

	for i, v := range z {
		fmt.Println("slice", i)
		for j, w := range v {
			fmt.Println("internal slice index", j, "internal slice value", w)
		}
	}
}
