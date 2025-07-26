package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Sleep()
	Walk()
}

type Human interface {
	Animal
	Speak()
	Think()
}

type Employee struct {
	Human
	Name string
	Age  int
}

type Cat struct {
	Name string
}

func main() {
	employee := &Employee{
		Name: "ali",
		Age:  30,
	}

	cat := &Cat{Name: "bob"}

	var human Human = employee
	var animal Animal = cat

	human.Eat()
	human.Sleep()
	human.Walk()
	human.Speak()
	human.Think()

	animal.Eat()
	animal.Sleep()
	animal.Walk()
}

func (emp *Employee) Eat() {
	fmt.Println("emp is eating\n")
}

func (emp *Employee) Sleep() {
	fmt.Println("emp is sleeping\n")
}

func (emp *Employee) Walk() {
	fmt.Println("emp is walking\n")
}

func (emp *Employee) Speak() {
	fmt.Println("Employee is speaking\n")
}

func (emp *Employee) Think() {
	fmt.Println("Employee is thinking\n")
}

func (cat *Cat) Eat() {
	fmt.Println("cat is eating\n")
}

func (cat *Cat) Sleep() {
	fmt.Println("cat is sleeping\n")
}

func (cat *Cat) Walk() {
	fmt.Println("cat is walking\n")
}
