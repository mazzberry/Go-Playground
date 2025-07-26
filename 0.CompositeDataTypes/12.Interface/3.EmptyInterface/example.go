package main

import (
	f "fmt"
)


type Person struct {
	Name string
}

func main() {
	print("hello")
	print(123)
	print(true)
	print([]int{123})
	print(map[string]int{"a" : 1, "b": 2})
	print(struct{Name string}{"Ali"})
}

func print(input interface{}) {
	switch input.(type) {
	case string:
		f.Println("string \n", input)
	case int:
		f.Println("int \n", input)
	case bool: 
		f.Println("bool \n", input)
	case []int:
		f.Println("slice \n", input)
	case map[string]int:
		f.Println("map \n", input)
	case Person:
		f.Println("struct \n", input)
	}
	f.Printf("%T : %v", input, input)     //any == emptyInterface or interface{}

} 


func AddOrder(order any) any { // any == interface{}
	if 1==2 {
		return 1
	}
	if order != 3 {
		return false
	} else {
		return true
	}
	
}