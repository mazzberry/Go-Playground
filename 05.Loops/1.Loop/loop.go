package main

import "fmt"


func main() {

	// 3 section loops
	// for initialization ; condition ; counter 

	sum := 0 
	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)


	// While loop
	a := 0

	for a < 10 {
		a++
	}
	fmt.Println(a)


	// infinite loop
	b := 0 

	for {
		b++
		//fmt.Println(b) // infinite
	}



}