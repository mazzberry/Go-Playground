package main

import "fmt"

func main() {
	msg := "database connection issue"
	sampleErr := fmt.Errorf("error occurred: %s", msg)
	fmt.Println(sampleErr)
}