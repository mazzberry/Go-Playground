package main

import "fmt"

type nonPositive struct {
	num int
}

func (e nonPositive) Error() string {
	return fmt.Sprintf("checkPositive: given number %d is not a positive number", e.num)
}

type notEven struct {
	num int
}

func (e notEven) Error() string {
	return fmt.Sprintf("checkEven: given number %d is not an even number", e.num)
}

func checkPositive(num int) error {
	if num <= 0 {
		return nonPositive{num: num}

	}
	return nil
}

func checkEven(num int) error {
	if num%2 != 0 {
		return notEven{num: num}
	}
	return nil
}


func checkPositiveAndEven(num int) error {
	if num > 100 {
		return fmt.Errorf("checkPositiveAndEven: given number %d is greater than 100", num)

	}

	err := checkPositive(num)
	if err != nil {
		return err
	}

	err = checkEven(num)
	if err != nil {
		return err
	}
	return nil
}




func main() {
	num := -3
	err := checkPositiveAndEven(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("checkPositiveAndEven: given number is positive and even")
	}
}