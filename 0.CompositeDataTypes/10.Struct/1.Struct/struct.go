package main



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