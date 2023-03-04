package parking

import (
	"errors"
	"fmt"
	"github.com/hbkkanna/parking/slot"
	"github.com/hbkkanna/parking/tariff"
	"time"
)

type Parkinglot interface {
	Park(vehicle slot.Vehicle) (slot.Ticket, error)
	UnPark(ticket slot.Ticket) (slot.Receipt, error)
}

type VehicleParkingLot struct {
	slots      map[int][]slot.Slot
	tariff     map[int]tariff.Tariff
	ticketCnt  int
	receiptCnt int
}

func (parkingLot *VehicleParkingLot) Park(vehicle slot.Vehicle) (slot.Ticket, error) {
	freeSlot, err := parkingLot.findFreeSlot(vehicle)
	if err != nil {
		return nil, err
	}
	freeSlot.SetInTime(time.Now())
	parkingLot.ticketCnt++
	return slot.NewTicket(parkingLot.ticketCnt, freeSlot), nil
}

func (parkingLot *VehicleParkingLot) UnPark(ticket slot.Ticket) (slot.Receipt, error) {
	vehicleSlot, err := parkingLot.findSlot(ticket.GetVehicleType(), ticket.GetNumber())
	if err != nil {
		return nil, err
	}
	err = ticket.SetOutTime(time.Now())
	if err != nil {
		return nil, err
	}
	cost := parkingLot.tariff[ticket.GetVehicleType()].GetCost(ticket)
	parkingLot.receiptCnt++
	receipt := slot.NewReceipt(parkingLot.receiptCnt, cost, slot.CloneVehicleSlot(vehicleSlot))
	vehicleSlot.Reset()
	return receipt, nil
}

func (parkingLot *VehicleParkingLot) findFreeSlot(vehicle slot.Vehicle) (slot.Slot, error) {
	slots, ok := parkingLot.slots[vehicle.GetVehicleType()]
	notAvail := errors.New(fmt.Sprintf(" No Slot for vehicle %v ", vehicle))
	if !ok {
		return nil, notAvail
	}
	for _, v := range slots {
		if v.IsFree() {
			return v, nil
		}
	}
	return nil, notAvail
}

func (parkingLot *VehicleParkingLot) findSlot(vehicleType int, number int) (slot.Slot, error) {
	slots, ok := parkingLot.slots[vehicleType]
	notFound := errors.New(fmt.Sprintf(" Slot not found for  vehicle type %d , for number %d ", vehicleType, number))
	if !ok {
		return nil, notFound
	}
	for _, v := range slots {
		if v.GetNumber() == number {
			return v, nil
		}
	}
	return nil, notFound
}

type ParkingConfig struct {
	vehicleType int
	slotCnt     int
	tariff      tariff.Tariff
}

func NewParkingConfig(vehicleType int, slotCnt int, tariff tariff.Tariff) *ParkingConfig {
	return &ParkingConfig{vehicleType: vehicleType, slotCnt: slotCnt, tariff: tariff}
}

func NewParkingLot(configs []*ParkingConfig) Parkinglot {
	slots := getSlotMap(configs)
	tariffs := getTariffMap(configs)
	return &VehicleParkingLot{
		slots:  slots,
		tariff: tariffs,
	}
}

func getSlotMap(configs []*ParkingConfig) map[int][]slot.Slot {
	vehicleSlots := make(map[int][]slot.Slot)
	for _, v := range configs {
		var slots []slot.Slot
		for j := 0; j < v.slotCnt; j++ {
			slots = append(slots, slot.NewVehicleSlot(slot.Vehicles[v.vehicleType], j))
		}
		vehicleSlots[v.vehicleType] = slots
	}
	return vehicleSlots
}

func getTariffMap(configs []*ParkingConfig) map[int]tariff.Tariff {
	tariffs := make(map[int]tariff.Tariff)
	for _, v := range configs {
		tariffs[v.vehicleType] = v.tariff
	}
	return tariffs
}
