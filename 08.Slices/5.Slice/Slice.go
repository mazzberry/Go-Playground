package main //append

import "fmt"


func main(){
	slice := []int{10, 20, 30, 40, 50}

	newSlice := slice[1:3]

	fmt.Println(len(newSlice)) // Print 2

	fmt.Println(cap(newSlice)) // Print 4

	newSlice = append(newSlice, 60)

	fmt.Println(len(newSlice)) // Print 3

	fmt.Println(cap(newSlice)) // Print 4



	slice2 := []int{10, 20, 30, 40, 50}

	newSlice2 := []int{}

	/*<<...>> or Ellipsis says that every element exists in slice2
	will append to newSlice2*/

	newSlice2 = append(newSlice2, slice...) 
	fmt.Println(slice2)
	fmt.Println(len(newSlice2)) // Print 5

	// remove an element from slice 

	slice3 := []int{10, 20, 30, 40, 50}
	slice3[2] = slice[len(slice3)-1]
	slice3 = slice3[:len(slice3)-1]
	fmt.Println(slice3) // Print [10 20 50 40]
	
	// remove an element by append function
	var slice4 = []int{1, 2, 3, 4, 5, 6}

	fmt.Println(removeElementByAppend(slice4, 2))

}


// remove an element by append function

func removeElementByAppend(slice []int, index int) []int {
	slice = append(slice[: index], slice[index+1:]...)
	return slice
}