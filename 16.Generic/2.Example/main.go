package main

import (
	"fmt"
)

type Number interface{
	int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64 | float32 | float64
}


func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%d\n", SumInt(mySlice))

	generic1 := []float64{1.3, 4.7, 6.6, 8.1, 9.2}
	fmt.Printf("%v\n", Sum(generic1))

	generic2 := []float64{3, 4, 7, 5, 2, 2, 1, 8}
	fmt.Printf("%v\n", Sum(generic2))

}

func SumInt(slc []int) (sum int) {
	for _, v := range slc {
		sum += v
	}
	return
}

func Sum[T Number](slc []T) (sum T) {
	for _, v := range slc {
		sum += v
	}
	return
}
