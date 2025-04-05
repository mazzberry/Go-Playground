package main // 2.2.7

import "fmt"

type employee struct {
	string
	Age    uint8
	Salary float32
}

/*
type employee struct {
	string  name
	int     age
	int     salary
	/ This one will cause an error, because two int fields are an anonymous in the struct
*/

func main() {
	// Create a new employee
	emp := employee{string: "Mohammad", Age: 21, Salary: 1000000.0}

	n := emp.string
	fmt.Printf("current name is: %s\n", n)
	// Change the name of the employee
	emp.string = "John Doe"

	n = emp.string
	fmt.Printf("new name is: %s\n", n)
}
