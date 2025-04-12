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
	fmt.Println("lion walks")

}
func main() {
	var a animal // تعریف یک متغیر از نوع interface
	a = lion{age: 5} // مقداردهی به متغیر از نوع struct
	a.breath() // فراخوانی متد breath
	a.walk() // فراخوانی متد walk

	a = &lion{age: 10} // مقداردهی به متغیر از نوع pointer به struct
	a.breath() // فراخوانی متد breath
	a.walk() // فراخوانی متد walk
}








