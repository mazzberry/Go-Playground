package main

import f "fmt"

func main() {

	defer f.Println("world") // execute last for defer keyWord
	f.Println("hello")       // first

}
