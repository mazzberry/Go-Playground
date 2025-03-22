package main 
import (
	"fmt"
)

func main () {
	arrayString := [3]string{"salam", "ali", "reza"}
	fmt.Printf("array %v, len %d cap %d", arrayString, len(arrayString), cap(arrayString))
}