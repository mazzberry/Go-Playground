package main

//import "fmt"

func main() {
	letters := []string{"ali", "hossein", "hassan"}

	for i := 1; i < 10; i++{
		// define a lable with name 'second' for this loop
		second:
			for i := 2; i < 9; i++ {
				for _, l := range letters {
					if l == "hossein" {
						// break the loop with second lable
						break second
					}
				}
			}
	}
}