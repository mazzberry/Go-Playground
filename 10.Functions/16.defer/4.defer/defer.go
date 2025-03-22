package main

import "fmt"

func main() {
	i := 1
	defer fmt.Println(i)
	i++
	fmt.Println(i)
	fmt.Println("First")
}
