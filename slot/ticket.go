package slot

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

func NewTicket(ticketNumber int, slot Slot) Ticket {
	return &VehicleTicket{
		Slot:         slot,
		ticketNumber: ticketNumber,
	}
}
