package main 

import "fmt"

// CRUD on Map
func main(){

	// Create
	animals := make(map[int]string)
	animals[1] = "Gopher"
	animals[2] = "owl"
	animals[3] = "cheetah"
	animals[4] = "lion"

	// Read
	fmt.Println("Read\n")
	fmt.Println(animals) 
	fmt.Println(animals[2])

	// Update
	fmt.Println("Update\n")
	fmt.Println(animals[2])
	animals[2] = "yooz"
	fmt.Println(animals[2])

	// Delete
	fmt.Println("Delete\n")
	fmt.Println(animals)
	fmt.Println(len(animals)) // 5  
	delete(animals, 4)  
	
	fmt.Println(animals)      // map[1:Gopher 2:owl 3:cheetah 5:lion]  
	fmt.Println(len(animals)) // 4

}