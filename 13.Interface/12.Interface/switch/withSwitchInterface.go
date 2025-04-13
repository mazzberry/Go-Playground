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
	fmt.Println("Lion breathes")
}

func (l lion) walk() {
	fmt.Println("Lion walks")
}

type dog struct {
	age uint8
}

func (d dog) breath() {
	fmt.Println("Dog breathes")
}

func (d dog) walk() {
	fmt.Println("Dog walks")
}

func main() {
	var a animal

	a = lion{age: 5}

	print(a)


}


func print(a animal) {
	
	switch v := a.(type) {

	case lion:
		fmt.Println("type: lion")

	case dog: 
		fmt.Println("type: dog")
	default:
		fmt.Printf("unknown type %T\n", v)
	}
}