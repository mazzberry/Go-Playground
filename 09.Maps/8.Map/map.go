package main

import "fmt"


func seriesStringToMap(inputs ...string) map[int]string{
	result := make(map[int]string)

	for index, input := range inputs {
		result[index] = input
	}

	return result
}


func mapToSlice(inputs map[int]string) []string{
	result := make([]string, len(inputs))

	for index, input := range inputs {
		result[index] = input
	}

	return result
}


func main() {

	myAnimal := "eagle, e, cheetah, lion, owl, gopher"
	
	myMappedAnimal := seriesStringToMap(myAnimal)
	fmt.Println(myMappedAnimal)

	mySliceAnimal := mapToSlice(myMappedAnimal)
	fmt.Println(mySliceAnimal)

	
}