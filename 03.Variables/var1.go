package main

import "fmt"

func main() {
	var s string = "Hello world" // for public variables we use this
	fmt.Printf(s, "\n")

	var a string
	a = "Hello world_2"
	fmt.Printf(a, "\n")

	b := "hello word 3 in short way" // this is only works in function (its not public variable)
	fmt.Printf(b, "\n")

	n, k, l := "hello", 1, 1.5 // define a few variables in one line
	var p, o, i = "world", 13, 24
	fmt.Printf(n, k, l, p, o, i, "\n")

	var y int
	var r float32
	var x bool
	var z string
	fmt.Printf("%v %v %v %q\n", y, r, x, z)

}
