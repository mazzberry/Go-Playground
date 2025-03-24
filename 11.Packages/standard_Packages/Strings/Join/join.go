package main

import (
	"strings"
	"fmt"
)

func main() {

	textSlice := []string{"Go", "will", "make", "you", "love", "programming", "again"}

	fmt.Println(strings.Join(textSlice, " "))

}