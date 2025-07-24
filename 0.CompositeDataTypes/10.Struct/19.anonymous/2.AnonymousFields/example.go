package main

import "fmt"

type Person struct {
	int
	string
	float64
}

func main() {
	person := Person{1, "Alireza", 1.76}
	fmt.Println(person)
}
