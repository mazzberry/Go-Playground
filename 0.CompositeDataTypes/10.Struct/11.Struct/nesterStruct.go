package main //2.2.8

import "fmt"

type Product struct {
	Name  string
	Price float64
}

type Mobile struct {
	Product  Product
	Ram      uint
	SimCount uint
}


func main() {
	var mobile Mobile = Mobile{}

	mobile.Product.Name = "iPhone 12"
	mobile.Product.Price = 1000.0
	mobile.Ram = 8
	mobile.SimCount = 2

	fmt.Printf("Mobile: %+v\n", mobile)

}
