package main

import (
	"errors"
	"fmt"
)

type errorOne struct{}
func (e errorOne) Error() string {
	return "Error One happend"
}

func main() {
	var err1 errorOne
	err2 := do() 
		if err1 == err2 {
			fmt.Println("equality operator: Both errors are equal")
		}
		if errors.Is(err2, err1) {
			fmt.Println("errors.Is: Both errors are equal")
		}
	}


func do() error {
	return errorOne{}

}

