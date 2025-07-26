package externalservices

import (
	"fmt"
	"notification/entities"
)

type EmailService struct {
}

func (e *EmailService) SendEmail(order *entities.Order) {
	fmt.Printf("Email sent : %v\n", order)
}

func NewEmailService() *EmailService {
	return &EmailService{}
}
