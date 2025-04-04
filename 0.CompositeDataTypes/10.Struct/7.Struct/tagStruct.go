package main // 2.2.6.1

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	Name   string  `json:"name"`
	Age    uint8   `json:"-"` // hidden field
	Salary float32 `json:"salary"`
}

func main() {
	emp := employee{Name: "MohammadReza", Age: 21, Salary: 15000.0}
	// Converting to JSON
	jsonData, err := json.MarshalIndent(emp, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(jsonData))
}
