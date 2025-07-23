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

type PersonBuilder struct {
	Person
}

func main() {

	//  1.
	person1 := Person{Name: "Maryam", Age: 20}

	person1.LastName = "Razzagh Nejad"

	// 2.
	personBuilder := &PersonBuilder{}
	person2 := personBuilder.SetName("Alireza").SetAge(19).SetLastName("Hosseini").SetGender("Male").SetHairColor("black").SetHeight(186).SetWeight(105)

	fmt.Printf("\nPerson1 : %+v\n", person1)
	fmt.Printf("\nPerson2 : %+v\n", person2)

}

func (builder *PersonBuilder) SetName(name string) *PersonBuilder {
	builder.Name = name
	return builder
}

func (builder *PersonBuilder) SetAge(age int) *PersonBuilder {
	builder.Age = age
	return builder
}

func (builder *PersonBuilder) SetLastName(lastName string) *PersonBuilder {
	builder.LastName = lastName
	return builder
}

func (builder *PersonBuilder) SetGender(gender string) *PersonBuilder {
	builder.Gender = gender
	return builder
}

func (builder *PersonBuilder) SetHeight(height float64) *PersonBuilder {
	builder.Height = height
	return builder
}

func (builder *PersonBuilder) SetWeight(weight float64) *PersonBuilder {
	builder.Weight = weight
	return builder
}

func (builder *PersonBuilder) SetHairColor(hairColor string) *PersonBuilder {
	builder.HairColor = hairColor
	return builder
}

func (builder *PersonBuilder) Build() Person {
	return builder.Person
}
