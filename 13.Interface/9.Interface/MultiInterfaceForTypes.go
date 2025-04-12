package main

import (
	"fmt"
	"unsafe"
)

type animal interface {
	breath()
	walk()
}

type mammal interface {
	feed()
}

type lion struct {
	age uint8
}

func (l lion) breath() {
	fmt.Println("Lion is breathing")
}

func (l lion) walk() {
	fmt.Println("lion walks")
}

func (l lion) feed() {
	fmt.Println("lion feeds young")
}

func main() {
	var a animal
	l := lion{}
	a = l
	a.breath()
	a.walk()
	var m mammal
	m = l
	m.feed()
	fmt.Println(unsafe.Sizeof(uint8(0)))
	fmt.Println(l.age)
}
