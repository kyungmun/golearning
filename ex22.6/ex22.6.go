package main

import (
	"fmt"
)

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)

	m[1] = Product{"kyungmun", 100}
	m[22] = Product{"연필", 1000}
	m[222] = Product{"지우개", 50}

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, 22)
	delete(m, 333)
	fmt.Println()

	for k, v := range m {
		fmt.Println(k, v)
	}
}
