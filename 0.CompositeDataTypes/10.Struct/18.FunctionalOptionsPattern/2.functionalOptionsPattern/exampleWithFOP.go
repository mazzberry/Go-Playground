package main

import "fmt"

type Person struct {
	Name      string
	Age       int
	LastName  string
	Gender    string
	Height    float64
	Weight    float64
	HairColor string
}

type PersonOptions func(*Person)  

func main() {

}

func SetName(name string) PersonOptions{
	return func(person *Person){
		person.Name = name
	}
}

func SetAge(age int) PersonOptions{
	return func(person *Person) {
		person.Age = age
	}
}

func SetLastName(lastName string) PersonOptions {
	return func(person *Person) {
		person.LastName = lastName
	}
}

func SetGender(gender string) PersonOptions{
	return func(person *Person) {
		person.Gender = gender
	}
}

func SetHeight(height float64) PersonOptions{
	return func(person *Person) {
		person.Height = height
	}
}

func SetWeight(weight float64) PersonOptions{
	return func(person *Person) {
		person.Weight = weight
	}
}

func SetHairColor(hairColor string) PersonOptions{
	return func(person *Person) {
		person.HairColor = hairColor
	}
}


