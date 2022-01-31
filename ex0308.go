package main

import "fmt"

func main() {

	favSport := ""

	switch favSport {
	case "football":
		{
			fmt.Println("football")
		}
	case "motogp":
		{
			fmt.Println("motogp")
		}
	default:
		{
			fmt.Println("no like cool sports")
		}
	}
}
