package main

import (
	"fmt"
)

func main() {
	a := 42

	pr(a)

	fmt.Println(a)

	prP(&a)

	fmt.Println(a)


	s := Student{
		Name: "Maziar",
	}
	fmt.Printf("s address: %p \n", &s)
	fmt.Println(s.GetterName())
	s.SetterName("mohsen")
	fmt.Println(s.GetterName())


}

func pr(a int) {
	a += 10
	fmt.Println(a)
}

func prP(a *int) {
	*a += 10
	fmt.Println(*a)
}

type Student struct {
	Name string
}

func (s Student) print() {
	fmt.Println(s.Name)
}

// getter
func (s Student) GetterName() string {
	fmt.Printf("s get: %p \n", &s)
	return s.Name
}

func (s *Student) SetterName(name string) {
	fmt.Printf("s setter address: %p \n", &s)
	s.Name = name
}