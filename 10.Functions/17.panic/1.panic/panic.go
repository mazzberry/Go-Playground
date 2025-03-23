package main

import "fmt"

func main() {

	a := []string{"a", "b"}
	printByIndex(a, 2)
}

func printByIndex(a []string, index int) {
	fmt.Println(a[index])
}