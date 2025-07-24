package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// API call and get respond

	apiResponse := struct {
		ResultCode        int
		ResultMsg         string
		TransactionAmount float64
		TransactionTime   string
	}{
		ResultCode:        0,
		ResultMsg:         "success",
		TransactionAmount: 100,
		TransactionTime:   "7/24/2025 - 4:50AM",
	}

	api_json, _ := json.Marshal(apiResponse)

	fmt.Printf(string(api_json))

}