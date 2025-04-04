package main

import (
    "fmt"
    "encoding/json"
	"log"
)

type employee struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	emp := employee{Name:"Mohammad",Age: 21,Salary: 15000.0}
	// Marshal function converts a struct to JSON format
	empJSON, err := json.Marshal(emp)
	if err != nil {
		log.Fatal(err.Error())
	}
	
	fmt.Printf("Marshal function output: %s\n", string(empJSON))

	// MarshallIndent function converts a struct to JSON format with indentation
	empJSON, err = json.MarshalIndent(emp, "", "   ")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("MarshalIndent function output: %s\n", string(empJSON))


}

