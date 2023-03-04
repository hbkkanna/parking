package slot

import (
	"errors"
	"fmt"
	"time"
)

type ParkingTime interface {
	GetInTime() time.Time
	SetInTime(time time.Time)
	GetOutTime() time.Time
	SetOutTime(time time.Time) error
	CalculateMinutes() float64
	CalculateHours() float64
}

type VehicleParkingTime struct {
	inTime  time.Time
	outTime time.Time
}

func (vehicleParkingTime *VehicleParkingTime) GetInTime() time.Time {
	return vehicleParkingTime.inTime
}

func (vehicleParkingTime *VehicleParkingTime) SetInTime(inTime time.Time) {
	vehicleParkingTime.inTime = inTime
}

func (vehicleParkingTime *VehicleParkingTime) GetOutTime() time.Time {
	return vehicleParkingTime.outTime
}

func (vehicleParkingTime *VehicleParkingTime) SetOutTime(outTime time.Time) error {
	if err := vehicleParkingTime.validateOutTime(outTime); err != nil {
		return err
	}
	vehicleParkingTime.outTime = outTime
	return nil
}

func (vehicleParkingTime *VehicleParkingTime) diffInOutTime() time.Duration {
	return vehicleParkingTime.outTime.Sub(vehicleParkingTime.inTime)
}

func (vehicleParkingTime *VehicleParkingTime) CalculateMinutes() float64 {
	return vehicleParkingTime.diffInOutTime().Minutes()
}

func (vehicleParkingTime *VehicleParkingTime) CalculateHours() float64 {
	return vehicleParkingTime.diffInOutTime().Hours()
}

func (vehicleParkingTime *VehicleParkingTime) validateOutTime(outTime time.Time) error {
	diff := outTime.Sub(vehicleParkingTime.inTime)
	if diff < 0 {
		return errors.New(fmt.Sprintf("invalid time in-time %v , out-time %v ", vehicleParkingTime.inTime, vehicleParkingTime.outTime))
	}
	return nil
}

func NewParkingTime() ParkingTime {
	return &VehicleParkingTime{}
}
