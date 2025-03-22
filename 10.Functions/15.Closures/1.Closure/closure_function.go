package main

import (
	"fmt"
)

func main() {
	number := 1

	func() {
		fmt.Println(number * 2)
	}()
}
