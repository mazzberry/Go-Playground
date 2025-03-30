package main

import "fmt"

var globalBlock = 125 // globalBlock is a global variable to this package

func main() {
	fmt.Println(globalBlock) // 125

	{
		var globalBlock = 200    // we overshadow the globalBlock with a local one in this block
		fmt.Println(globalBlock) // 200
	}

	{
		//var x = 14 // x is local to this block
	}

}
