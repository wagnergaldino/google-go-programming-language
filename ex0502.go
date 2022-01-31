package main

import "fmt"

type person struct {
	firstName, lastName string
	iceCreamFlavours    []string
}

func main() {
	p1 := person{
		firstName:        "Wagner",
		lastName:         "Galdino",
		iceCreamFlavours: []string{"peanut", "corn"},
	}
	p2 := person{
		firstName:        "Marcelo",
		lastName:         "Monaco",
		iceCreamFlavours: []string{"peanut", "corn"},
	}

	m := map[string]person{
		`Galdino`: p1,
		`Monaco`:  p2,
	}

	for i, v := range m {
		fmt.Println("map key", i)
		fmt.Println(v.firstName, v.lastName)
		for _, w := range v.iceCreamFlavours {
			fmt.Println(w)
		}
	}

}
