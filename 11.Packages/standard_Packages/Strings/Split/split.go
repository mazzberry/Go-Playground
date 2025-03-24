package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Go will make you love programming again. We promise"

	textSlice := strings.Split(text, " ")

	for _, item := range textSlice {
		fmt.Println(item)
	}
}