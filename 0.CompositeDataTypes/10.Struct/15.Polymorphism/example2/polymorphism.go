package main

import "fmt"

type BusTicket struct {
	Id                int
	DepartureCity     string
	ArrivalCity       string
	DepartureTime     string
	BusType           string
	PassengerName     string
	DepartureTerminal string
	ArrivalTerminal   string
	NationalCode      string
	Price             int
}

type FlightTicket struct {
	Id               int
	DepartureCity    string
	ArrivalCity      string
	DepartureTime    string
	AirplaneType     string
	PassengerName    string
	DepartureAirport string
	ArrivalAirport   string
	ArrivalTime      string
	PassportId       string
	PassengerType    string
	Price            int
}

func main() {
	busTicket := BusTicket{
		Id:                1,
		BusType:           "Scania",
		DepartureCity:     "Tehran",
		ArrivalCity:       "Mashhad",
		DepartureTime:     "10:00",
		DepartureTerminal: "Terminal shargh",
		ArrivalTerminal:   "Imam Reza Bus Suburban Terminal",
		NationalCode:      "4222212212",
		PassengerName:     "Alireza",
		Price:             50000,
	}

	flightTicket := FlightTicket{
		Id:               1,
		DepartureCity:    "izmir",
		ArrivalCity:      "New york",
		DepartureTime:    "10:00",
		DepartureAirport: "Turkey",
		ArrivalAirport:   "New york Airport",
		ArrivalTime:      "13:00",
		AirplaneType:     "Sub",
		PassengerName:    "MohammadReza",
		PassportId:       "4220522212",
		Price:            110000,
	}

	printer := []TicketPrinter{busTicket, flightTicket}

	for _, printer := range printer {
		printer.PrintTicket()
	}

}

type TicketPrinter interface {
	PrintTicket()
}

func (ticket BusTicket) PrintTicket() {
	fmt.Printf("bus Ticket:\n  ID: %d\n Departure City: %s\n Arrival City: %s\n Departure Time: %s\n Bus Type: %s\n Passenger Name: %s\n Departure Terminal: %s\n Arrival Terminal: %s\n National Code: %s\n Price: %d\n", ticket.Id, ticket.DepartureCity, ticket.ArrivalCity, ticket.DepartureTime, ticket.BusType, ticket.PassengerName, ticket.DepartureTerminal, ticket.ArrivalTerminal, ticket.NationalCode, ticket.Price)

}

func (ticket FlightTicket) PrintTicket() {
	fmt.Printf("\nFlight Ticket:\n  ID: %d\n Departure City: %s\n Arrival City: %s\n Departure Time: %s\n Airplane Type: %s\n PassengerName: %s\n Departure Airport: %s\n Arrival Airport: %s\n Passport Id: %s\n Price: %d\n", ticket.Id, ticket.DepartureCity, ticket.ArrivalCity, ticket.DepartureTime, ticket.AirplaneType, ticket.PassengerName, ticket.DepartureAirport, ticket.ArrivalAirport, ticket.PassportId, ticket.Price)

}
