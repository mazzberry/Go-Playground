package main // Bad Practice of Abstraction

import "fmt"

const (
	BaseSalary       = 5600000
	ExtraPerHourRate = 90000
	HourlyWorkRate   = 110000
	TaxRate          = 0.09
	shiftSalaryRate  = 80000
)

type Employee struct {
	Id           int
	NationalCode string
	FullName     string
	Type         string // "full" or "part"
	Hours        float64
}

func main() {
	employees := []Employee{
		{Id: 1, NationalCode: "1234567890", FullName: "Ali Rezaei", Type: "FullTime", Hours: 8},
		{Id: 2, NationalCode: "0987654321", FullName: "Sara Hosseini", Type: "PartTime", Hours: 4},
		{Id: 3, NationalCode: "1122334455", FullName: "Mohammad Karimi", Type: "FullTime", Hours: 10},
		{Id: 4, NationalCode: "5566778899", FullName: "Fatemeh Ahmadi", Type: "Shift", Hours: 120},
	}
	for _, employee := range employees {
		salary, tax := employee.SalaryCalculator(float64(employee.Hours))
		fmt.Printf("Employee ID: %d, Name: %s, National Code: %s, Type: %s, Salary: %.2f, Tax: %.2f\n", employee.Id, employee.FullName, employee.NationalCode, employee.Type, salary, tax)
	}
}

func (employee Employee) FullSalaryCalculator(hours float64) (salary float64, tax float64) {
	salary = BaseSalary + (ExtraPerHourRate*HourlyWorkRate)*1.4
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee Employee) PartSalaryCalculator(hours float64) (salary float64, tax float64) {
	salary = BaseSalary + (ExtraPerHourRate*HourlyWorkRate)*1.4
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee Employee) ShiftSalaryCalculator(hours float64) (salary float64, tax float64) {
	salary = shiftSalaryRate * hours
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee Employee) SalaryCalculator(hours float64) (salary float64, tax float64) {
	if employee.Type == "FullTime" {
		return employee.FullSalaryCalculator(hours)
	} else if employee.Type == "PartTime" {
		return employee.PartSalaryCalculator(hours)
	} else {
		return employee.ShiftSalaryCalculator(hours)
	}
}
