package main

import "fmt"

func addCount(x *int) {
	*x += 1
	fmt.Printf("value = %d, address in memory = %p\n", *x, x)
}

func addCountWithoutPointer(x int) {

	x += 1
	fmt.Printf("value = %d, address in memory = %p\n", x, &x)
}

func printCount(x int) {
	fmt.Printf("value = %d, address in memory = %p\n", x, &x)
}

func main() {
	var count int

	addCount(&count)
	addCount(&count)

	addCountWithoutPointer(count)
	fmt.Printf("value = %d, address in memory = %p\n", count, &count)

	printCount(count)
}
