package main

import (
	"fmt"
)


func main(){
	
	slice := make([]int, 5) // تعریف یک اسلایس با اندازه مشخص 
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice2 := make([]int, 3, 5) // با ظرفیت و اندازه مشخص
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

}

// مقدار ظرفیت هرگز نباید بیشتر از اندازه باشد.

