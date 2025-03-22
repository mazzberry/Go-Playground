package main

import (
	"fmt"
)

func main() {

	var sum = func(n1, n2 int) int {
		sum := n1 + n2

		return sum
	}

	result := sum(5, 7)

	fmt.Println("sum is: ", result)

}