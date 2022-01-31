package main

import "fmt"

func main() {

	func() {
		fmt.Println("função anonima sem parâmetro")
	}()

	func(x int) {
		fmt.Println("função anonima com parâmetro", x)
	}(7)

}
