package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Go will make you love programming again. We promise"

	fmt.Println("HasPrefix Go:", strings.HasPrefix(text, "Go"))
	fmt.Println("HasPrefix se:", strings.HasPrefix(text, "se"))

	fmt.Println("HasSuffix se:", strings.HasSuffix(text, "se"))
	fmt.Println("HasSuffix Go:", strings.HasSuffix(text, "Go"))
}