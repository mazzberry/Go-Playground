package main

import	"fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 5)

	numberOfElementCopied := copy(dst, src)
	fmt.Println(numberOfElementCopied, dst)
	
}