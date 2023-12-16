package main

import (
	"fmt"
	"os"
)

type opfunc func(a, b int) int

func getOperator(op string) opfunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}

	} else {
		return nil
	}
}

func main() {
	fn := getOperator("/")
	fn2 := getOperator("*")

	if fn != nil {
		result := fn(1, 2)
		fmt.Println(result)
	}
	fmt.Println(fn2(3, 5))
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return

	}
	defer f.Close()
	fmt.Fprintln(f, "kyungmun test")
}
