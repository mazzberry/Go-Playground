package main // 2.6.3

import (
	"errors"
	"fmt"
)

func main() {
	sampleErr := errors.New("error occurred")
	fmt.Println(sampleErr)
}

