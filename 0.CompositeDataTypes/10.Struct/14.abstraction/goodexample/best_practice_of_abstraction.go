package main

import "fmt"

func main() {
	FullTimeEmployees := []FullTimeEmployee{
		{Id: 1, NationalCode: "1234567890", FullName: "Ali Rezaei", Type: "FullTime", ExtraHours: 40},
		{Id: 3, NationalCode: "1122334455", FullName: "Mohammad Karimi", Type: "FullTime", ExtraHours: 110},
	}

	PartTimeEmployees := []PartTimeEmployee{
		{Id: 2, NationalCode: "0987654321", FullName: "Sara Hosseini", PartTimeHours: 20},
		{Id: 1, NationalCode: "242351141", FullName: "Hamid", PartTimeHours: 40},
		{Id: 4, NationalCode: "3519864", FullName: "Hasan", PartTimeHours: 30},
	}

	// ShiftEmployee := []ShiftEmployee{
	// 	{Id: 1, NationalCode: "13512124125", FullName: "homeira", ShiftHours: 120},
	// 	{Id: 2, NationalCode: "1234567890", FullName: "Ali Rezaei", ShiftHours: 80},
	// 	{Id: 3, NationalCode: "1312521232", FullName: "Mohammad", ShiftHours: 140},
	// }

	//1
	var employee Employee = FullTimeEmployees[0]
	salary, tax := employee.SalaryCalculator(float64(FullTimeEmployees[0].ExtraHours))
	fmt.Printf("Employee: %d, Salary: %.2f, Tax: %.2f\n", employee, salary, tax)
	//2
	employee = FullTimeEmployees[1]
	salary, tax = employee.SalaryCalculator(float64(FullTimeEmployees[1].ExtraHours))
	fmt.Printf("Employee: %d, Salary: %.2f, Tax: %.2f\n", employee, salary, tax)

	//3
	employee = PartTimeEmployees[0]
	salary, tax = employee.SalaryCalculator(float64(PartTimeEmployees[0].PartTimeHours))
	fmt.Printf("Employee: %d, Salary: %.2f, Tax: %.2f\n", employee, salary, tax)
	//4
	employee = PartTimeEmployees[1]
	salary, tax = employee.SalaryCalculator(float64(PartTimeEmployees[1].PartTimeHours))
	fmt.Printf("Employee: %d, Salary: %.2f, Tax: %.2f\n", employee, salary, tax)

}

type Employee interface {
	SalaryCalculator(hours float64) (salary float64, tax float64)
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
	Type         string
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

func (employee FullTimeEmployee) SalaryCalculator(hours float64) (salary float64, tax float64) {
	salary = BaseSalary + (ExtraPerHourRate*HourlyWorkRate)*1.4
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee PartTimeEmployee) SalaryCalculator(hours float64) (salary float64, tax float64) {
	salary = BaseSalary + (ExtraPerHourRate*HourlyWorkRate)*1.4
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee ShiftEmployee) SalaryCalculator(hours float64) (salary float64, tax float64) {
	salary = shiftSalaryRate * hours
	tax = salary * TaxRate
	salary += tax
	return
}
