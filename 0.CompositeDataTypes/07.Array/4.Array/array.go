package main

import (
	"fmt"
)


func main() {
	
	nums := [2][2][2]int{{{1, 2}, {2, 3}}, {{4, 5}, {6, 7}}}
	fmt.Printf("array nums value: %v, len: %d, cap: %d", nums, len(nums), cap(nums))

}