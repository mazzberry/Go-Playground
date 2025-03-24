package main

import "fmt"

func main() {

	a := []string{"a", "b"}
	checkAndPrint(a, 2)
	fmt.Println("Existing normally")

}

func checkAndPrint(s []string, index int) {

	defer handelOutBounds()

	if index > (len(s) - 1) {
		panic("Out of bound access for slice")

	}

	fmt.Println(s[index])

}

func handelOutBounds() {
	if r := recover(); r != nil {
		fmt.Println("recovering from panic :", r)
	}

}
