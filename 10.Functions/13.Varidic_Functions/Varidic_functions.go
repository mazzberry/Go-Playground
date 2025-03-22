package main 

import(
	"fmt"
)

func sum (nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2, 3, 4, 5)
	sum(1,8,9)

	nums := []int{1,19,2}  // define an array 
	sum(nums...) // if wanna pass an array to varidic function as argument 
}

/* 
در اینجا تابعی تعریف کردیم که تعداد دلخواهی از آرگومان‌ها از نوع 
int
را به کمک 
… (بهش میگن Ellipsis) 
که قبل از نوع داده قرار گرفته به داخل تابع منتقل می‌کند.
برای صدا زدن این توابع می‌توان به روش 
sum(num1, num2, …)
 عمل کرد.
اگر شما داده‌ای با نوع 
slice 
دارید می‌توانید آن را به کمک اپراتور 
…(Ellipsis) 
به صورت 
sum(nums…) 
به داخل تابع انتقال بدید.
*/