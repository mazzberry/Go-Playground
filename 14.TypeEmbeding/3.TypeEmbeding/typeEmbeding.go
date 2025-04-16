package main

import "fmt"

type animal interface {
	breath()
	walk()
}

type dog struct {
	age uint8
}

func (d dog) breath() {
	fmt.Println("dog breathes")
}

func (d dog) walk() {
	fmt.Println("dog walks")
}

type pet1 struct {
	a animal
	name string
}

type pet2 struct {
	animal
	name string
}


func main() {
	d := dog{age: 5}
	p1 := pet1{name: "Milo", a: d}

	fmt.Println(p1.name)
    // p1.breathe()
    // p1.walk()
	p1.a.breath()
	p1.a.walk()
	
	p2 := pet2{name: "Oscar", animal: d}
	fmt.Println(p2.name)
	p2.breath()
	p2.walk()
	p1.a.walk()
	p1.a.breath()



}








