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
	fmt.Println("lion walks")
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


func callBreath(a animal) {
	a.breath()
}

func callWalk(a animal) {
	a.walk()
}


func main() {
	l := lion{age: 10}
	callBreath(l)
	callWalk(l)

	d := dog{age: 5}
	callBreath(d)
	callWalk(d)

}
