package main

import "fmt"

func main() {
	a := 20
	b := &a
	c := &b

	fmt.Println("pointer b value: ", b)
	fmt.Println("pointer c value: ", c)
	fmt.Println("pointer c point to: ", *c)
	fmt.Println("get value by double pointer: ", **c)
}