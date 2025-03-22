package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c"}

	//With index and value
	fmt.Println("Both Index and Value")
	for i, letter := range letters {
		fmt.Printf("Index: %d Value:%s\n", i, letter)
	}

	//Only value
	fmt.Println("\nOnly value")
	for _, letter := range letters {
		fmt.Printf("Value: %s\n", letter)
	}

	//get value by index
	fmt.Println("\nGet value by index")
	for i := range letters {
		fmt.Printf("letters: %s\n", letters[i])	}

}
