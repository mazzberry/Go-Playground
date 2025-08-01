package main

import (
	"fmt"

	jalaali "github.com/jalaali/go-jalaali"                                      //outer module
	"github.com/mazzberry/Go-Playground/tree/main/14.Modules/2.example/services" //inner package

)

func main() {
	fmt.Printf("Hello world!\n")

	year, month, day, err := jalaali.ToGregorian(1401, 15, 29)
	if err == nil {
		fmt.Printf("%d %d %d\n", year, month, day)
	} else {
		fmt.Printf("Error: %s\n", err)
	}

	var service services.TestService = services.TestService{}

	var service2 services.TestService2 = services.TestService2{}

	fmt.Printf("%v", service)
	fmt.Printf("%v", service2)

	

}
