package main

import "fmt"

func main() {
	var dayOfWeek int

	fmt.Printf("Enter a day of week: ")
	fmt.Scanln(&dayOfWeek)

	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
		fallthrough
	case 2:
		fmt.Println("Tuesday")
		fallthrough
	case 3:
		fmt.Println("Wednesday")
		fallthrough
	case 4:
		fmt.Println("Thursday")
		fallthrough
	case 5:
		fmt.Println("Friday")
		fallthrough
	case 6:
		fmt.Println("Saturday")
		fallthrough
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Day")
	}
}