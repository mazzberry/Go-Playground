package main

import "fmt"


func main() {
	// Arithmetic op
	a := 5
	b := 10

	c := a + b
	d := a - b
	e := a * b
	f := a / b
	g := a % b

	println(c, d, e, f, g, "\n")

	println("Comparative op\n")

	//Comparative op
	h := 5
	i := 10
	m := 5

	println(h == i) // false
  	println(h == m) // true
 	println(h < i)  // true
  	println(i > m)  // trume
  	println(m <= h) // true
  	println(m >= i) // false
  	println(i != h) // true
  	println(m!= h) // false

	println("\nLogical op\n")
	
	//Logical op

	l := true
	k := true 
	j := false

	fmt.Println(l && k)
	fmt.Println(l || k)
	fmt.Println(!l)
	fmt.Println(!j)


}
