package main

import "fmt"

func add10AndSum(num1 int, num2 int, sum func(n1, n2 int) int) {
	result := sum(num1+10, num2+10)
	fmt.Println("Sum by adding 10 is:", result)
}

func main() {
	add10AndSum(5, 3, func(n1, n2 int) int {
		sum := n1 + n2

		return sum
	})
}

/*
 ما یک تابع تعریف کردیم که در پارامتر سوم یک تابع دریافت می‌کند که باید دو ورودی
 int
 و یک خروجی
 int
 داشته باشد.
*/
