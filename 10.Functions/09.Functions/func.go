package main

import (
	"fmt"
)

func plus(a int, b int) int { // or func plus(a, b int){} if they're the same type
	return a + b
}

func multiply(t, l int) int {
	return t * l
}

func name(v string, i, p int) {
	fmt.Println(v, i, p)
}

func main() {
	fmt.Println(plus(4, 13))
	fmt.Println(multiply(10, 3))
	name("Hello", 1, 2)
}
