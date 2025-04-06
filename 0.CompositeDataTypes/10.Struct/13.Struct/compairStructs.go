package main

import "fmt"

type employee struct {
	Name   string
	Age    uint8
	Salary float32
}

func main() {
	employee1 := employee{
		Name:   "Maryam",
		Age:    20,
		Salary: 12000.0,
	}
	employee2 := employee{
		Name:   "Mohammadreza",
		Age:    21,
		Salary: 13000.0,
	}

	// Compare the two structs
	if employee1 == employee2 {
		fmt.Println("employee1 and employee2 are equal")
	} else {
		fmt.Println("employee1 and employee2 are not equal")
	}
}
