package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println("number =", i, "- reminder(4) =", i%4)
		} else if i%3 == 0 {
			fmt.Println("number =", i, "- reminder(3) =", i%3)
		} else {
			fmt.Println("number =", i, "- reminder(2) =", i%2)
		}
	}
}
