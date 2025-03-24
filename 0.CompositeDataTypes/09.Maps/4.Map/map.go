package main 

import "fmt"


func main(){

	var personData = map[string]string{
		"Name": "John Doe",
		"family": "Doe",
		"dob": "01-01-1990",
		"city": "New York",
	}	

	name, nameExist := personData["Name"]
	age, ageExist := personData["age"]
	dob, dobExist := personData["dob"]
	city, cityExist := personData["city"]
	organization, organizationExist := personData["organization"]
	
	fmt.Println(name, nameExist)
	fmt.Println(age, ageExist)
	fmt.Println(dob, dobExist)
	fmt.Println(city, cityExist)
	fmt.Println(organization, organizationExist)
	


}