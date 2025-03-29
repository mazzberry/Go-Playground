package main

import "fmt"


var glob = 125 // glob is global to this package
var glob2 = 250 // glob2 is global to this package

func main() {
	fmt.Println(glob) // 125
	var glob2 = 300 // glob2 is local to this main function and shadows glob2
	fmt.Println(glob2)

	{
		//var x = 14 // x is local to this block
	}
	
	//fmt.Println(x) // 14 // Error: x is not defined in this scope

}