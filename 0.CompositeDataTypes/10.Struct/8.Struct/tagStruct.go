package main

import (
	"encoding/json"
	"log"
)

type employee struct {
	Name   string  `json:"name,omitempty"`
	Age    uint8   `json:"age,omitempty"`
	Salary float32 `json:"salary,omitempty"`
}

func main() {
	emp := employee{Name: "John Doe", Age: 30,} // create an instance of employee without giving salary value
	jsonData, err := json.MarshalIndent(emp, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("JSON data: %s\n", jsonData) // Output: JSON data: {"name":"John Doe","age":30}
}
