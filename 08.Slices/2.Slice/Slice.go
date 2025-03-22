// 1.8.4.3
package main

import "fmt"

func main(){
	slice := []string{"red", "blue", "green", "yellow"}

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	intSlice := []int {10, 20, 30}

	fmt.Println(cap(intSlice))
	fmt.Println(cap(intSlice))
	

}