package main

import (
	"fmt"
	"runtime/debug"
)

func main() {

	a := []string{"a", "b"}
	checkAndPrint(a, 2)

}

func checkAndPrint(s []string, index int) {
	defer handelOutBounds()

	if index > (len(s) - 1) {
		panic("out of bound access for slice")
	}
	fmt.Println(s[index])

}

func handelOutBounds() {

	if r := recover(); r != nil {
		fmt.Println("Recovering form panic :", r)
		fmt.Println("stack trace :")
		debug.PrintStack()
	}

}
