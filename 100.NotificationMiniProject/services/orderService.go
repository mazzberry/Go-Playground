package services

import (
	"fmt"
	"notification/entities"
	"notification/externalservices"
)

type OrderService struct {
	EmailService *externalservices.EmailService
	SmsService   *externalservices.SmsService
}

func (o *OrderService) CreateOrder(order *entities.Order) *entities.Order {
	fmt.Printf("Order created : %v", order)

	o.SmsService.SendMessage(order)
	o.EmailService.SendEmail(order)

	return order
}

func NewOrderService() *OrderService {
	return &OrderService{
		EmailService: externalservices.NewEmailService(),
		SmsService:   externalservices.NewSmsService(),
	}
}
