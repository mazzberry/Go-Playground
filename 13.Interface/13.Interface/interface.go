package main

import "fmt"

func main() {
	test("this is string")
	test("10")
	test(true)

	test2("this is string")
	test2("10")
	test2(true)
}

func test(a interface{}) { // any is an alias for interface{}
	fmt.Printf("(%v, %T) \n", a, a)
}

func test2(a any) { // any is an alias for interface{}
	fmt.Printf("(%v, %T) \n", a, a)
}
