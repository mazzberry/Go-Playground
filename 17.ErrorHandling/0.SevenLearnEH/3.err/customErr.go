package main

import (
	"fmt"
	"net/http"

	"io"
)

type HttpError struct {
	StatusCode int
	Message    string
}

func main() {
	response, err := GetReq("") //http://dummyjson3123.com/products/categories
	if err != nil {
		fmt.Println(err)
	return
	}
	fmt.Println(response)

}

func (e HttpError) Error() string {  // CREATED CUSTOM ERROR
	return fmt.Sprintf("Http error occurred status code :%d msg: %s", e.StatusCode, e.Message)
}

func NewHttpError(statusCode int, message string) *HttpError { 
	return &HttpError{StatusCode: statusCode, Message: message}
}

func GetReq(url string) (string, error) {
	if len(url) == 0{
		return "", NewHttpError(400, "error occurred (url is empty)")
	}
	response, err := http.Get(url)
	if err != nil {
		return "", NewHttpError(500, "error occurred")
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return "", NewHttpError(400, "error occurred (body is empty)")
	}
	return string(responseBody), nil

}