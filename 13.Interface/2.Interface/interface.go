package main

import "fmt"

type animal interface {
	breath()
	walk()
}

type lion struct {
	age uint8
}

func (l lion) breath() {
	fmt.Println("Lion breath")
}

func (l lion) walk() {
	fmt.Println("Lion walk")
}

func main() {
	var a animal
	a = lion{age: 5}
	a.breath()
	a.walk()
	fmt.Println(a) // 
}
