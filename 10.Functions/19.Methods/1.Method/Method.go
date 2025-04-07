package main

import "fmt"

type employee struct {
	name   string
	age    uint8
	salary float32
}

func (e employee) getterName() string {
	return e.name
}
func (e employee) getterAge() uint8 {
	return e.age
}

func (e employee) getterSalary() float32 {
	return e.salary
}

func main() {
	emp1 := employee{
		name:   "Maryam",
		age:    20,
		salary: 12000.0,
	}
	fmt.Println("Name:", emp1.getterName())
	fmt.Println("Age:", emp1.getterAge())
	fmt.Println("Salary:", emp1.getterSalary())
}
