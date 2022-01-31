package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["q"] = []string{`Being boss`, `banana`, `dildo`}

	for i, v := range m {
		fmt.Println("map key", i)
		for j, w := range v {
			fmt.Println("map slice index", j, "map slice value", w)
		}
	}
}
