package main // 2.6.5

import (
	"os"
	"fmt"
)

func main() {
	file, _ := os.Open("nonexistent.txt")
	fmt.Println(file)
}