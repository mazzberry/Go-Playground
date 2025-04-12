package main

import "fmt"

type animal interface {
	breath()
	walk()
}

type lion struct {
	age uint8
}

func (l *lion) breath() {
	fmt.Println("lion breaths")

}

func (l *lion) walk() {
	fmt.Println("lion walks")
}

func main() {
	var a animal // تعریف یک متغیر از نوع interface

	
	a = &lion{age: 5}
	a.breath() // فراخوانی متد breath
	a.walk()   // فراخوانی متد walk
}
/*	a = lion{age: 10} we'll face an error here because lion is a pointer and we are trying to assign a value to it.
	a.breath() */
