package main

import "fmt"

func main() {

	sample := map[string]string{
		"a": "x",
		"b": "y",
	}

	// Iterating over all key and values -> key:value
	fmt.Println("both key and value \n")

	for k, v := range sample {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}

	// Iterating over only keys
	fmt.Println("only keys\n")
	for k := range sample {
		fmt.Printf("key: %s\n", k)
	}

	// Iterating over only value
	fmt.Println("only values")
	for _, v := range sample {
		fmt.Printf("value: %s\n", v)
	}

	//for-range loop on string
	sample2 := "a£b"

    //With index and value
    fmt.Println("Both Index and Value")
    for i, letter := range sample2 {
        fmt.Printf("Start Index: %d Value:%s\n", i, string(letter))
    }

    //Only value
    fmt.Println("\nOnly value")
    for _, letter := range sample2 {
        fmt.Printf("Value:%s\n", string(letter))
    }

    //Only index
    fmt.Println("\nOnly Index")
    for i := range sample2 {
        fmt.Printf("Start Index: %d\n", i)
    }
}
