package main

import "fmt"

func main() {
	sum := 0
	for {
		sum++
		if sum == 10 {
			break
		}
	}

	fmt.Println(sum)
	fmt.Println("now this line will execute")
}
