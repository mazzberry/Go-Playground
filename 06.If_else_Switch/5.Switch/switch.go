package main

import "fmt"

func main() {

	switch ch := "d"; ch {

	case "a":
		fmt.Println("a")

	case "b", "c":
		fmt.Println("b or c")

	default:
		fmt.Println("no matching character")

	}

}
