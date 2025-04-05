package main // 2.2.8

import "fmt"

type employee struct {
	Name    string
	Age     int
	Salary  float64
	address address
}

type address struct {
	Street string
	City   string
	State string
}


func main() {
	m_Adrress := address{Street: "hamase", City: "Tehran", State: "Iran"}
	emp := employee{Name: "Mohammad", Age: 21, Salary: 15000.0, address: m_Adrress}
	fmt.Printf("city: %s\n", emp.address.City)
	fmt.Printf("state: %s\n", emp.address.State)
	
}