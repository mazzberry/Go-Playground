package main

import "fmt"

type employee struct {
	name    string
	age     uint8
	salary  float32
	address address
}

type address struct {
	country string
	state   string
	city    string
}

func (a address) getterDetails() {
	fmt.Println("Country:", a.country)
	fmt.Println("State:", a.state)
	fmt.Println("City:", a.city)
}

func main() {
	emp1_address := address{
		country: "Iran",
		state:   "Tehran",
		city:    "Tehran",
	}
	emp1 := employee{
		name:    "Mohammadreza",
		age:     21,
		salary:  12000.0,
		address: emp1_address,
	}

	emp1.address.getterDetails()

}
