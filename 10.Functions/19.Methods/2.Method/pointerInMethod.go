package main

import "fmt"

type employee struct {
	name   string
	age    uint8
	salary float32
}

func (e *employee) setterName(newName string) {
	e.name = newName

}

func main() {
	emp1 := employee{
		name:   "Ali",
		age:    20,
		salary: 12000.0,
	}

	emp1.setterName("Maryam")
	fmt.Println("Name:", emp1.name)
}
