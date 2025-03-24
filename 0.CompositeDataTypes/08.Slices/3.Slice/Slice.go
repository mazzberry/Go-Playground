package main

import "fmt"

func main() {

	slice := []int{99: 88}

	fmt.Println(slice)


    // declare an empty Slice

	sliceOne := make([]int, 0)

	sliceTwo := []int{}

	fmt.Println(sliceOne == nil) // print false

	fmt.Println(len(sliceOne)) // print 0

	fmt.Println(cap(sliceOne)) // print 0

	fmt.Println(sliceTwo == nil) // print false

	fmt.Println(len(sliceTwo)) // print 0

	fmt.Println(cap(sliceTwo)) // print 0


	// revaluing a Slice

	var vSlice = []int{10, 20, 30, 40}
	
	fmt.Println(vSlice)

	vSlice[1] = 25

	fmt.Println(vSlice)


	// Create a new slice based on a predefined slice
	
	var x = []int{10, 20, 30, 40,  50}
	
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	var y = x[1:3]

	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))


	
}
