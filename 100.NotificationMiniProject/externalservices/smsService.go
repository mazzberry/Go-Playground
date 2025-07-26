package externalservices

import (
	"fmt"
	"notification/entities"
)

type SmsService struct {
}

func (e *SmsService) SendMessage(order *entities.Order) {
	fmt.Printf("Sms sent : %v\n", order)
}

func NewSmsService() *SmsService {
	return &SmsService{}
}
