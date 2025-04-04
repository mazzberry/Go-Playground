package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	emp := employee{Name: "John Doe", Age: 30, Salary: 50000.0}
	// Converting to json
	empJSON, err := json.MarshalIndent(emp, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(empJSON))

	// if they say: the first letter of Name field should be small letter
	// we can use the following code to convert it to lower case

	type lowerEmployee struct {
		Name   string  `json:"name"`
		Age    int     `json:"age"`
		Salary float64 `json:"salary"`
	}

	lemp := lowerEmployee{Name: "ali", Age: 25, Salary: 60000.0}
	lempJSON, err := json.MarshalIndent(lemp, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(lempJSON)) 
	/*
	if any field in the struct starts with a lower case letter, it will be privet and won't be included in the json output.
	*/

}
