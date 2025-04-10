package main

import "fmt"

type taxSystem interface {
	calculateTax() int
}

type indianTax struct {
	taxPercentage int
	income        int
}

func (i *indianTax) calculateTax() int {
	tax := (i.income * i.taxPercentage) / 100
	return tax
}

type singaporeTax struct {
	taxPercentage int
	income        int
}

func (s *singaporeTax) calculateTax() int {
	tax := (s.income * s.taxPercentage) / 100
	return tax
}

type usaTax struct {
	taxPercentage int
	income        int
}

func (u *usaTax) calculateTax() int {
	tax := (u.income * u.taxPercentage) / 100
	return tax
}

func main() {
	indianTax := &indianTax{taxPercentage: 10,
		income: 100000,
	}
	singaporeTax := &singaporeTax{
		taxPercentage: 15,
		income:        200000,
	}

	taxSystem := []taxSystem{indianTax, singaporeTax}
	totalTax := calculateTotalTax(taxSystem)
	fmt.Println("Total Tax: ", totalTax)
}

func calculateTotalTax(taxSystems []taxSystem) int {
	var totalTax int
	for _, t := range taxSystems {
		totalTax += t.calculateTax() // در اینجا runtime polymorphism رخ می دهد
	}
	return totalTax

}
