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

type dog struct {
	age uint8
}

func (d dog) breath() {
	fmt.Println("Dog breath")
}

func (d dog) walk() {
	fmt.Println("Dog walk")
}

func main() {
	var a animal


	a = lion {age: 5}
	print(a)

}

func print(a animal) {
	l, ok := a.(lion)
	if ok {
		fmt.Printf("age : %d\n", l.age)
	}
}
 



