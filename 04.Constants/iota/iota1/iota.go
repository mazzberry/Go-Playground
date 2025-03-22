package main

import (
	"fmt"
)

const (
	a = 0
	b = 1
	c = 2
)

const (
	d = iota
	e // 1
	f // 2
)

// creating enum by iota

type Size uint8

const (
	small      Size = iota
	medium          // 1
	large           // 2
	extralarge      // 3
)

// ignoring the first value
const (
	_      = iota
	red    // 1
	blue   // 2
	yellow // 3

)

func main() {
	fmt.Println(a, b, c, "\n", d, e, f)

	fmt.Println(small)
	fmt.Println(medium)
	fmt.Println(large)
	fmt.Println(extralarge)

	fmt.Println(red, blue, yellow)
}
