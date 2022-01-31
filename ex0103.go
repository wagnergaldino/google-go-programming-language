package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s, _ := fmt.Printf("%d - %s - %t", x, y, z)
	fmt.Println(s)
}
