package main

import "fmt"

type animal interface {
	breath()
	walk()
}

func main() {
	var a animal
	fmt.Println(a) // <nil>
}
