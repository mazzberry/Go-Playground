package main 

import "fmt"

type animal interface {
	breath()
	walk()
}


type cat string 

func (c cat) breath() {
	fmt.Println("cat breath")
}


func (c cat) walk() {
	fmt.Println("cat walks")
}


func main() {
	var a animal 

	a = cat("smokey")
	a.breath()
	a.walk()
}



