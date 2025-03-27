package main

import "fmt"

func changeNumberByReference(num *int) int {
	*num = *num + 2

	return *num // Returning the address of the modified variable, not the value itself
}

func changeNumberByValue(num int) int {
	num = num + 2

	return num // Returning the address of the modified variable, not the value itself
}

func main() {

	i := 20

	// Change number by reference
	fmt.Println("addres of i : ", &i)
	fmt.Println(changeNumberByReference(&i))
	fmt.Println("value of i after change by reference : ", i)

	// Change number by value
	fmt.Println("addres of i : ", &i)
	fmt.Println(changeNumberByValue(i))
	fmt.Println("value of i after change by value : ", i)

}
