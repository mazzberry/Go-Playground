package main

import "fmt"


func main() {

	var totalSum int = 846
	var number int = 19
	var avg float32

	avg = float32(totalSum) / float32(number)

	fmt.Printf("Average = %v\n", avg) // %f instead of %v
}