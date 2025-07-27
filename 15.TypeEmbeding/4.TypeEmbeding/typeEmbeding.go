package main

import "fmt"

type animal interface {
	breath()
	walk()
}

type human interface {
	animal
	speak()
}

type employee struct {
	name string
}

func (e employee) breath() {
	fmt.Println("employee breathes")
}

func (e employee) walk() {
	fmt.Println("employee walks")
}

func (e employee) speak() {
	fmt.Println("employee speaks")
}


func main() {
	var h human 

	h = employee{name: "John"}
	h.breath()
	h.walk()
	h.speak()
}