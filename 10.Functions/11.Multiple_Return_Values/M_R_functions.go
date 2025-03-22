package main

import (
	"fmt"
)

func vals() (int, int) {
	return 3, 7
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)


}

/*
در کد بالا تابعی ساختیم با اسم 
vals 
که دو خروجی از نوع 
int 
دارد بنابراین نوع تعریف توابع چندبازگشتی متفاوت است و بصورت 
(data-type, data-type, … ) 
است.
در مثال اول ما با کمک دو متغیر a,b دو خروجی تابع 
vals 
را دریافت کردیم.
در مثال دوم ما با استفاده از 
_ blank identifier 
از دریافت یا استخراج خروجی اول صرف نظر کردیم و فقط خروجی دوم با متغیر 
c 
را دریافت کردیم.
*/