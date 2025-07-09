package main // polymorphism

import "fmt"

func main() {
	FullTimeEmployees := []FullTimeEmployee{
		{Id: 1, NationalCode: "1234567890", FullName: "Ali Rezaei", ExtraHours: 40},
		{Id: 3, NationalCode: "1122334455", FullName: "Mohammad Karimi", ExtraHours: 110},
	}

	PartTimeEmployees := []PartTimeEmployee{
		{Id: 2, NationalCode: "0987654321", FullName: "Sara Hosseini", PartTimeHours: 20},
		{Id: 1, NationalCode: "242351141", FullName: "Hamid", PartTimeHours: 40},
		{Id: 4, NationalCode: "3519864", FullName: "Hasan", PartTimeHours: 30},
	}

	ShiftEmployee := []ShiftEmployee{
		{Id: 1, NationalCode: "13512124125", FullName: "homeira", ShiftHours: 120},
		{Id: 2, NationalCode: "1234567890", FullName: "Ali Rezaei", ShiftHours: 80},
		{Id: 3, NationalCode: "1312521232", FullName: "Mohammad", ShiftHours: 140},
	}

	var employees []Employee = []Employee{}

	for _, employee := range FullTimeEmployees {
		employees = append(employees, employee)
	}

	for _, employee := range PartTimeEmployees {
		employees = append(employees, employee)
	}

	for _, employee := range ShiftEmployee {
		employees = append(employees, employee)
	}

	for _, employee := range employees {
		salary, tax := employee.SalaryCalculator()
		fmt.Printf("\nEmployee (%T): %v\n: salary: %f\n, tax %f\n", employee, employee, salary, tax)
	}
}

type Employee interface {
	SalaryCalculator() (salary float64, tax float64)
}

const (
	BaseSalary       = 5600000
	ExtraPerHourRate = 90000
	HourlyWorkRate   = 110000
	TaxRate          = 0.09
	shiftSalaryRate  = 80000
)

type FullTimeEmployee struct {
	Id           int
	NationalCode string
	FullName     string
	ExtraHours   float64
}

type PartTimeEmployee struct {
	Id            int
	NationalCode  string
	FullName      string
	PartTimeHours float64
}

type ShiftEmployee struct {
	Id           int
	NationalCode string
	FullName     string
	ShiftHours   float64
}

func (employee FullTimeEmployee) SalaryCalculator() (salary float64, tax float64) {
	salary = BaseSalary + (ExtraPerHourRate*HourlyWorkRate)*1.4
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee PartTimeEmployee) SalaryCalculator() (salary float64, tax float64) {
	salary = HourlyWorkRate * employee.PartTimeHours
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee ShiftEmployee) SalaryCalculator() (salary float64, tax float64) {
	salary = shiftSalaryRate * employee.ShiftHours
	tax = salary * TaxRate
	salary += tax
	return
}
