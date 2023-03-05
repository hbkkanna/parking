package slot

import (
	"fmt"
)

type Ticket interface {
	Slot
	GetTicketNumber() int
}

type VehicleTicket struct {
	Slot
	ticketNumber int
}

func (vehicleTicket *VehicleTicket) GetTicketNumber() int {
	return vehicleTicket.ticketNumber
}

func (vehicleTicket *VehicleTicket) String() string {
	return fmt.Sprintf("Parking Ticket: \n  Ticket Number: %d \n  Spot Number: %d \n  "+
		"Entry Date-Time: %v ", vehicleTicket.GetTicketNumber(), vehicleTicket.GetNumber(), vehicleTicket.GetInTime())
}

func NewTicket(ticketNumber int, slot Slot) Ticket {
	return &VehicleTicket{
		Slot:         slot,
		ticketNumber: ticketNumber,
	}
}
