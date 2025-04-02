package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary float32
}

//add: implement employee struct with pointer demonstration
func main() {
	
	emp := employee{name: "Mohammad", age: 21, salary: 15000.0}
	empP := &emp // pointer to the struct
	fmt.Printf("Emp: %+v\n", empP)
	empP = &employee{name: "John Doe", age: 30, salary: 1000.0} // change the value of the struct through the pointer
	fmt.Printf("Emp: %+v\n", empP)


}
