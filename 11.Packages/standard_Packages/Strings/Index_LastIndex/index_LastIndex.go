package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "Go will make you love programming again. We promise"

	fmt.Println(strings.Index(text, "o"))
	fmt.Println(strings.LastIndex(text, "o"))
}