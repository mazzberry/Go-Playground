package main // 2.6.1

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("non_existent_file.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	} else {
		fmt.Println(file.Name() + "File opened successfully")
	}
}

