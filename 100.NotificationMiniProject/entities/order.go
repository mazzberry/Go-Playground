package entities

type Order struct {
	Id int
	UserFullName string
	UserEmail string
	UserPhone string
	Price float64
	Status bool
}