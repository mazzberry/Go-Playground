package main

import (
	"errors"
	"fmt"
)

func main() {
	output, err := createErrorMethod1(2)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)

	// fmt.Println(createErrorMethod1(0))
	// fmt.Println(createErrorMethod2(0))

	// fmt.Println(createErrorMethod1(10))
	// fmt.Println(createErrorMethod1(10))

}

func createErrorMethod1(number int) (int, error) {
	if number == 0 {
		return 0, errors.New("M1: number is not valid")
	}
	
	return number * 6, nil
}

func createErrorMethod2(number int) (int, error) {
	if number == 0 {
		return 0, fmt.Errorf("M2: number is not valid: %d", number)
	}
	
	return number * 6, nil
}
