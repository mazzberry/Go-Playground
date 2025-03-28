package main

import "fmt"

func main() {
	// 1
	var s string = "Hello world" // for public variables we use this
	fmt.Printf(s, "\n")
	// 2
	var a string
	a = "Hello world_2"
	fmt.Println(a)
	// 3
	var i3 = 3
	var f1 = 3.44
	var s1 = "Hello world" // for public
	fmt.Printf("%v, %T\n", i3, i3)
	fmt.Printf("%v, %T\n", f1, f1)
	fmt.Printf("%v, %T\n", s1, s1)
	// 4
	b := "hello word 3 in short way" // this is only works in function (its not public variable)
	fmt.Println(b)
	// 5
	n, k, l := "hello", 1, 1.5 // define a few variables in one line
	var p, o, i = "world", 13, 24
	fmt.Println(n, k, l, p, o, i)
	// 6

	var (
		i7 = 12
		f5 = 3.78
		s2 = "hello world" // for public variables we use this
	)
	fmt.Printf("%v, %T\n", i7, i7)
	fmt.Printf("%v, %T\n", f5, f5)
	fmt.Printf("%v, %T\n", s2, s2)

	var y int
	var r float32
	var x bool
	var z string
	fmt.Printf("%v %v %v %q\n", y, r, x, z)

}
