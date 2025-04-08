package main

import "fmt"

type animal interface {
	breath() string
	walk() string
}

func main() {
	var a animal
	fmt.Println(a) // <nil>  because: default value of interface is nil
	// fmt.Println(a.breath()) // this will not work because a is nil

}

