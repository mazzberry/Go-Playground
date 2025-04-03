package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func main() {
	empP := new(employee)
	fmt.Printf("Emp Pointer Address: %p\n", empP)
	fmt.Printf("Emp Pointer: %+v\n", empP)
	fmt.Printf("Emp Value: %+v\n", *empP)
	*empP = employee{"John Doe", 30, 1000}
	fmt.Printf("Emp After Change: %+v\n", *empP) // modify the value of the struct through the pointer
}
