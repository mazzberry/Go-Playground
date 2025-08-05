package main

import (
	"errors"
	"fmt"
)

func main() {
	x, err := createErrorMethod1(0)
	fmt.Println(x, err)

	x, err = createErrorMethod2(0)
	fmt.Println(x, err)
}

func createErrorMethod1(number int) (int, error){
	if number == 0 {
		return 0, errors.New("M1: number is not valid")
	}
	err := errors.New("and error has occurred")
	fmt.Println(err)
	return number * 6, nil
}

func createErrorMethod2(number int) (int, error){
	if number == 0 {
		return 0, fmt.Errorf("M2: number is not valid: %d", number)
	}
	err := errors.New("and error has occurred")
	fmt.Println(err)
	return number * 6, nil
}