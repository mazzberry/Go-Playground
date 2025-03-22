package main

import (
	"fmt"
)

const s string = "hello world!"

func main() {
	fmt.Println(s)

	// defining constant without needing declare type

	// named untype constant
	const a = 1
	const b = "circle"
	const c = 5.4
	const d = true
	const e = 'a'
	const f = 3 + 5i

	var u = 123      //Default hidden type is int
	var v = "circle" //Default hidden type is string
	var w = 5.6      //Default hidden type is float64
	var x = true     //Default hidden type is bool
	var y = 'a'      //Default hidden type is rune
	var z = 3 + 5i   //Default hidden type is complex128

	fmt.Println("")
	fmt.Printf("Type: %T Value: %v\n", a, a)
	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %v\n", c, c)
	fmt.Printf("Type: %T Value: %v\n", d, d)
	fmt.Printf("Type: %T Value: %v\n", e, e)
	fmt.Printf("Type: %T Value: %v\n", f, f)

	fmt.Printf("Type: %T Value: %v\n", u, u)
	fmt.Printf("Type: %T Value: %v\n", v, v)
	fmt.Printf("Type: %T Value: %v\n", w, w)
	fmt.Printf("Type: %T Value: %v\n", x, x)
	fmt.Printf("Type: %T Value: %v\n", y, y)
	fmt.Printf("Type: %T Value: %v\n", z, z)

}
