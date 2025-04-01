package main

import "fmt"

func main() {
	a := 20
	b := &a
	c := &b

	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", b)
	fmt.Printf("c = %d\n", c)

	fmt.Println()
	fmt.Printf("a = %d\n", a)
	fmt.Printf("*&a = %d\n", *&a)
	fmt.Printf("*b = %d\n", *b)
	fmt.Printf("**c = %d\n", **c)

	fmt.Println()
	fmt.Printf("&a: %d\n", &a)
	fmt.Printf("b: %d\n", b)
	fmt.Printf("&*b: %d\n", &*b)
	fmt.Printf("*&b: %d\n", *&b)
	fmt.Printf("*c: %d\n", *c)

	fmt.Println()
	fmt.Printf("&b: %d\n", &b)
	fmt.Printf("c: %d\n", c)
	fmt.Printf("*c: %d\n", *c)
	fmt.Printf("**c: %d\n", **c)

}