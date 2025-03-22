package main

import (
	"fmt"
)

func main() {
	animals := make(map[int][]string) // nil map of string-int pairs
	animals[0] = []string{"Gopher", "running", "rodent"}
	animals[1] = []string{"owl", "flying", "carnivorous"}
	animals[2] = []string{"cheetah", "running", "carnivorous"}
	animals[3] = []string{"eagle", "flying", "carnivorous"}
	animals[4] = []string{"lion", "running", "carnivorous"}

	for index, animal := range animals {
		fmt.Printf("%v- %s is %s animal and can %s \n", index, animal[0], animal[2], animal[1])
	}
}