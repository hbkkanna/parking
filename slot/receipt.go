package slot

type Receipt interface {
	Slot
	GetReceiptNumber() int
	GetCost() float64
}

type VehicleReceipt struct {
	Slot
	cost          float64
	receiptNumber int
}

func (vehicleReceipt *VehicleReceipt) GetReceiptNumber() int {
	return vehicleReceipt.receiptNumber
}

func (vehicleReceipt *VehicleReceipt) GetCost() float64 {
	return vehicleReceipt.cost
}

func NewReceipt(receiptNumber int, cost float64, slot Slot) Receipt {
	return &VehicleReceipt{
		Slot:          slot,
		receiptNumber: receiptNumber,
		cost:          cost,
	}
}
