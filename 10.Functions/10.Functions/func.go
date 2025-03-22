package main

import "fmt"

func main() {
	plus := func(a int, b int) int {
		return a + b
	}

	fmt.Println(plus(3, 4))
}
