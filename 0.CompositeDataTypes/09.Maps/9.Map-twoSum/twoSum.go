package main 

func twoSum(nums []int, target int) []int {
	keys := map[int]int{}
	for i, item := range nums{
		compliment := target - item

		if _, ok := keys[compliment]; ok {
			return []int{keys[compliment], i}
	}
	keys[item] = i

	}
	return []int{0, 0}
}

func main () {
	nums := []int{2, 7, 11, 15}
	target := 18
	result := twoSum(nums, target)
	println(result[0], result[1])
}