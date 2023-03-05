package slot

import (
	"fmt"
)

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

func (vehicleReceipt *VehicleReceipt) String() string {
	return fmt.Sprintf("Parking Receipt: \n  Receipt Number: R-%d \n  "+
		"Entry Date-Time: %v \n  Exit Date-Time: %v \n  Cost: %.2f",
		vehicleReceipt.GetReceiptNumber(), vehicleReceipt.GetInTime(),
		vehicleReceipt.GetOutTime(), vehicleReceipt.cost)
}

func NewReceipt(receiptNumber int, cost float64, slot Slot) Receipt {
	return &VehicleReceipt{
		Slot:          slot,
		receiptNumber: receiptNumber,
		cost:          cost,
	}
}
