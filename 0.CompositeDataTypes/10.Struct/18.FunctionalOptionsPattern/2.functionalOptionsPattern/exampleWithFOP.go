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
	person := NewPerson(SetName("Mohammadreza"), SetLastName("Hosseini"), SetAge(21), SetGender("Male"), SetHeight(184), SetWeight(120), SetHairColor("Brown"))

	fmt.Printf("Person : %+v", person)
}

func NewPerson(options ...PersonOptions) *Person {
	person := &Person{Name: "Alireza", Age: 19}

	for _, option := range options {
		option(person)
	}
	return person
}

func SetName(lastName string) PersonOptions {
	return func(person *Person) {
		person.Name = lastName
	}
}

func SetAge(age int) PersonOptions {
	return func(person *Person) {
		person.Age = age
	}
}

func SetLastName(lastName string) PersonOptions {
	return func(person *Person) {
		person.LastName = lastName
	}
}

func SetGender(gender string) PersonOptions {
	return func(person *Person) {
		person.Gender = gender
	}
}

func SetHeight(height float64) PersonOptions {
	return func(person *Person) {
		person.Height = height
	}
}

func SetWeight(weight float64) PersonOptions {
	return func(person *Person) {
		person.Weight = weight
	}
}

func SetHairColor(hairColor string) PersonOptions {
	return func(person *Person) {
		person.HairColor = hairColor
	}
}
