package main

import (
	"fmt"
	"io"

	"net/http"
)

func main() {
	response, err := http.Get("https://www.gogle.com")

	if err != nil {
		fmt.Println("an error has occurred on get request")
		return
	}
	fmt.Println(response.Status)

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("an error has occurred on reading the response body")
	}

	fmt.Println(string(responseBody))
}
