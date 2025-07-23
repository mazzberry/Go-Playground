package main //best practice of Abstraction

import "fmt"

const (
	BaseSalary       = 5600000
	ExtraPerHourRate = 90000
	HourlyWorkRate   = 110000
	TaxRate          = 0.09
	shiftSalaryRate  = 80000
)

func main() {
	FullTimeEmployees := []FullTimeEmployee{
		{Employee: Employee{Id: 1, NationalCode: "1234567890", FullName: "Ali Rezaei"}, Type: "FullTime", ExtraHours: 40},
		{Employee: Employee{Id: 3, NationalCode: "1122334455", FullName: "Mohammad Karimi"}, Type: "FullTime", ExtraHours: 110},
	}

	PartTimeEmployees := []PartTimeEmployee{
		{Employee: Employee{Id: 2, NationalCode: "0987654321", FullName: "Sara Hosseini"}, PartTimeHours: 20},
		{Employee: Employee{Id: 1, NationalCode: "242351141", FullName: "Hamid"}, PartTimeHours: 40},
		{Employee: Employee{Id: 4, NationalCode: "3519864", FullName: "Hasan"}, PartTimeHours: 30},
	}

	ShiftEmployee := []ShiftEmployee{
		{Employee: Employee{Id: 1, NationalCode: "13512124125", FullName: "homeira"}, ShiftHours: 120},
		{Employee: Employee{Id: 2, NationalCode: "1234567890", FullName: "Ali Rezaei"}, ShiftHours: 80},
		{Employee: Employee{Id: 3, NationalCode: "1312521232", FullName: "Mohammad"}, ShiftHours: 140},
	}

	//1
	var employee EmployeeSalaryCalculator = FullTimeEmployees[0]
	salary, tax := employee.SalaryCalculate(float64(FullTimeEmployees[0].ExtraHours))
	fmt.Printf("\nfullEmployee: %v, Salary: %.2f, Tax: %.2f\n", employee, salary, tax)
	//2
	employee = FullTimeEmployees[1]
	salary, tax = employee.SalaryCalculate(float64(FullTimeEmployees[1].ExtraHours))
	fmt.Printf("\nfullEmployee: %v, Salary: %.2f, Tax: %.2f\n", employee, salary, tax)

	//3
	employee = PartTimeEmployees[0]
	salary, tax = employee.SalaryCalculate(float64(PartTimeEmployees[0].PartTimeHours))
	fmt.Printf("\npartEmployee: %v, Salary: %.2f, Tax: %.2f\n", employee, salary, tax)
	//4
	employee = PartTimeEmployees[1]
	salary, tax = employee.SalaryCalculate(float64(PartTimeEmployees[1].PartTimeHours))
	fmt.Printf("\npartEmployee: %v, Salary: %.2f, Tax: %.2f\n", employee, salary, tax)
	//5
	employee = ShiftEmployee[0]
	salary, tax = employee.SalaryCalculate(float64(ShiftEmployee[0].ShiftHours))
	fmt.Printf("\nshiftEmployee: %v, Salary: %.2f, Tax: %.2f\n", employee, salary, tax)
	//6
	employee = ShiftEmployee[1]
	salary, tax = employee.SalaryCalculate(float64(ShiftEmployee[1].ShiftHours))
	fmt.Printf("\nshiftEmployee: %v, Salary: %.2f, Tax: %.2f\n", employee, salary, tax)

}

type EmployeeSalaryCalculator interface {
	SalaryCalculate(hours float64) (salary float64, tax float64)
}

type Employee struct {
	Id           int
	NationalCode string
	FullName     string
	Unit         string
}

type FullTimeEmployee struct {
	Employee
	Type       string
	ExtraHours float64
}

type PartTimeEmployee struct {
	Employee
	PartTimeHours float64
}

type ShiftEmployee struct {
	Employee
	ShiftHours float64
}

func (employee FullTimeEmployee) SalaryCalculate(hours float64) (salary float64, tax float64) {
	salary = BaseSalary + (ExtraPerHourRate*HourlyWorkRate)*1.4
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee PartTimeEmployee) SalaryCalculate(hours float64) (salary float64, tax float64) {
	salary = BaseSalary + (ExtraPerHourRate*HourlyWorkRate)*1.4
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee ShiftEmployee) SalaryCalculate(hours float64) (salary float64, tax float64) {
	salary = shiftSalaryRate * hours
	tax = salary * TaxRate
	salary += tax
	return
}
