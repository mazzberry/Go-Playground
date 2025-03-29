package main

import "fmt"


var glob = 125 // glob is global to this package

func main() {

	{
		//var x = 14 // x is local to this block
	}
	fmt.Println(glob) // 125
	//fmt.Println(x) // 14 // Error: x is not defined in this scope

}