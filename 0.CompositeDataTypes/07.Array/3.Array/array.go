package main

import (
	"fmt"
)


func main() {
	nums := [...]int {1, 25,30, 45, 892, 78, 541, 11}
	fmt.Printf("array nums value %v, len %d, cap %d", nums, len(nums), cap(nums))
}