package main

import (
	"fmt"
)

func main() {

		// [inclusive, exclusive]
	// or we can say ==> [from, up to]

	k := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	l := k[:]   // slice of all elements
	m := k[3:]  // slice from 4th element to end
	n := k[:6]  // slice first 6 elements
	o := k[3:6] // slice the 4th, 5th and 6th elements

	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)

	// pointing to the same data

	c := [...]int{1, 2, 3}
	d := &c

	d[1] = 5

	fmt.Println(c)
	fmt.Println(d)
}
