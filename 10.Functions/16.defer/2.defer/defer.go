package main

import f "fmt"

func main() {

	// anonymous function with defer
	defer func() { f.Println("in inline defer") }()
	f.Println("executed")

}
