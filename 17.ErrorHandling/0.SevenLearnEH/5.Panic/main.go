package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	function1(nums, 7) // it'll panic

	Divide(10, 2)

	Divide(10, 0)
}

func function1(numbers []int, index int) {
	// 1- index out of range
	// fmt.Println(numbers[index])

	// 2- panic
	// panic("test") //causes panic

	// 3- log
	// log.Fatal("Fatal") //causes panic
}

func Divide(a, b int) {
	defer func() {
		fmt.Printf("defer of divide: %d , %d  ", a, b)
		if err := recover(); err != nil{
			fmt.Println("Error: ", err)
			fmt.Println("recovering from err ...") //it
		}
	}()
	fmt.Printf("start of divide: %d , %d\n", a, b)
	fmt.Printf("Result : %d\n", a/b)
	fmt.Printf("end of divide: %d , %d\n", a, b)
}
