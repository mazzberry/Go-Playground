package main

import (
	"fmt"

	jalaali "github.com/jalaali/go-jalaali"
)

func main() {
	fmt.Printf("Hello world!\n")

	year, month, day, err := jalaali.ToGregorian(1397, 1, 1)
	if err == nil {
		fmt.Printf("%d %d %d\n", year, month, day)
	} else {
		fmt.Printf("Error: %s\n", err)
	}
}

