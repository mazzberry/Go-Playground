package main

import "fmt"

type employee struct { // define a struct type for employee
	name   string
	age    int
	salary float32
}

var emp = employee{} // create an instance of employee

var emp1 = employee{name: "John Doe", age: 30, salary: 5000.0} // create an instance of employee with given values

var emp2 = employee{ // create an instance of employee with given values using pointer
	name: "Jane Doe",
 	age: 28,
  	salary: 6000.0,
  } 


  func main() {
	emp3 := new(employee) // create a pointer to an instance of employee
	emp3.name = "ali" // it's not nessecary to initialize age in this case, as it's zero value by default.
	emp3.salary = 7000.0
	fmt.Printf("Employee 4: %+v\n", emp3)
  }