package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Go will make you love programming again. We promise"

	textSlice := strings.Split(text, " ")
	// s := make([]string, 0)

	for _, item := range textSlice {
		// s = append(s, item)
		fmt.Println(item)
	}
	// fmt.Println(s)
}
