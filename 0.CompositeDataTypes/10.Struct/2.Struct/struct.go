package main

import "fmt"

type employee struct { // define a struct type for employee
	name   string
	age    int
	salary float32
}

func main() {

	var emp = employee{} // create an instance of employee

	fmt.Printf("Employee 1: %+v\n", emp)

	var emp1 = employee{name: "John Doe", age: 30, salary: 5000.0} // create an instance of employee with given values
	fmt.Printf("Employee 2: %+v\n", emp1)

	// create an instance of employee with given values using pointer
	var emp2 = employee{
		name:   "Jane Doe",
		age:    28,
		salary: 6000.0,
	}
	fmt.Printf("Employee 3: %v\n", emp2) // if use %+v instead of %v, it will print the key-value pairs of the struct

	emp3 := employee{name: "ali", salary: 7000.0} // it's not nessecary to initialize age in this case, as it's zero value by default.
	fmt.Printf("Employee 4: %+v\n", emp3)

	fmt.Println(emp1.name, emp2.name)

	emp1.name = "john smith" // update the value of name field {{2.2.3}}

	fmt.Printf("Employee 1 after update: %+v\n", emp1)

}
