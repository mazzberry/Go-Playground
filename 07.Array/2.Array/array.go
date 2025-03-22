package main 
import (
	"fmt"
)

func main () {
	nums := [5]int {}
	fmt.Printf("array nums values %v, len %d, cap %d", nums, len(nums), cap(nums))

	nums[0] = 1
	nums[1] = 2
	nums[2] = 10
	nums[4] = 999

	fmt.Println("")
	fmt.Printf("array nums values %v, len %d, cap %d", nums, len(nums), cap(nums))

}