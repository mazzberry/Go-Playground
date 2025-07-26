// main --> orderService -> emailService || smsService
package main

import (
	"notification/entities"
	"notification/services"
)

func main() {
	order := entities.Order{
		Id:           1,
		UserFullName: "John Doe",
		UserEmail:    "john@gmail.com",
		UserPhone:    "9114642212",
		Price:        2000,
		Status:       true,
	}

	orderService := services.NewOrderService()
	orderService.CreateOrder(&order)

}
