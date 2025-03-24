package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "Golang go Go will make you love programming again. We promise"

	fmt.Println(strings.Count(text, "Go"))

}
