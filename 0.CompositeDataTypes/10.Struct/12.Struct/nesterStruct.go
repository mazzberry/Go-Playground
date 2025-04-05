package main //2.2.8

import "fmt"

type Product struct {
	Name  string
	Price float64
}

type Mobile struct {
	Product
	Ram      uint
	SimCount uint
}


func main() {
	var mobile Mobile = Mobile{}

	mobile.Name = "iPhone 12"
	mobile.Price = 1000.0
	mobile.Ram = 8
	mobile.SimCount = 2

	fmt.Printf("Mobile: %+v\n", mobile)

}
