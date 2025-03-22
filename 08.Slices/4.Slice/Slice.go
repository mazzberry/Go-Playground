package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}
	slice1 := slice[1]
	var slice3 = slice[3]

	newSlice := slice[1:3] // index out of range

	newSlice[3] = 45

	fmt.Println(slice1)
	fmt.Println(slice3)

	fmt.Println(newSlice) // err index out of range

}
