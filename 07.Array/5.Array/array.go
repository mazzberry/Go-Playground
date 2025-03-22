package main

import (
	"fmt"
)

func main() {

	nums := [2]int{1, 2}
	nums1 := [2]int{1, 3}
	nums2 := [2]int{1, 2}
	//nums3 := [3]int{1, 2, 3}
	//chars := [2]string{"a", "b"}

	fmt.Println(nums == nums1) // false
	fmt.Println(nums == nums2)	// true
	//fmt.Println(nums == nums3) // error: invalid operation: nums == nums4 (mismatched types [2]int and [3]int)

	//fmt.Println(nums == chars) // error: invalid operation: nums == chars (mismatched types [2]int and [2]str
}
