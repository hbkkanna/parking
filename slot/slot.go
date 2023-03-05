package slot

import (
	"time"
)

const (
	SCOOTER = iota
	SUV
	TRUCK
)

var Vehicles map[int]Vehicle

func init() {
	Vehicles = map[int]Vehicle{SCOOTER: NewRoadVehicle(SCOOTER), SUV: NewRoadVehicle(SUV), TRUCK: NewRoadVehicle(SUV)}
}

type Vehicle interface {
	GetVehicleType() int
}

type RoadVehicle struct {
	vehicleType int
}

func (vehicle *RoadVehicle) GetVehicleType() int {
	return vehicle.vehicleType
}

type Slot interface {
	Vehicle
	ParkingTime
	GetNumber() int
	Reset()
	IsFree() bool
}

type VehicleSlot struct {
	Vehicle
	ParkingTime
	number int
}

func (vehicleSlot *VehicleSlot) GetNumber() int {
	return vehicleSlot.number
}

func (vehicleSlot *VehicleSlot) IsFree() bool {
	zeroVal := time.Time{}
	return vehicleSlot.GetInTime() == zeroVal
}

func (vehicleSlot *VehicleSlot) Reset() {
	vehicleSlot.SetOutTime(time.Time{})
	vehicleSlot.SetInTime(time.Time{})
}

func NewVehicleSlot(vehicle Vehicle, number int) Slot {
	return &VehicleSlot{
		Vehicle:     vehicle,
		number:      number,
		ParkingTime: NewParkingTime(),
	}
}

func CloneVehicleSlot(vehicleSlot Slot) Slot {
	clonedSlot := NewVehicleSlot(NewRoadVehicle(vehicleSlot.GetVehicleType()), vehicleSlot.GetNumber())
	clonedSlot.SetInTime(vehicleSlot.GetInTime())
	clonedSlot.SetOutTime(vehicleSlot.GetOutTime())
	return clonedSlot
}

func NewRoadVehicle(vehicleType int) Vehicle {
	return &RoadVehicle{
		vehicleType: vehicleType,
	}
}
