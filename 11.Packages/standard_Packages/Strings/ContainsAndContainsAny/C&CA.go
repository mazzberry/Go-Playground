package main 

import (
	"fmt"
	"strings"
)


func main() {
	text := "go will make you love programming again. I promise."

	fmt.Println(strings.Contains(text, "love"))
	fmt.Println(strings.Contains(text, "Golang"))
	fmt.Println(strings.ContainsAny(text, "Golang"))
}



