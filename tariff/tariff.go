package tariff

import (
	"github.com/hbkkanna/parking/slot"
	"math"
)

type ModelCalculator interface {
	GetCost(parkingTime slot.ParkingTime) float64
}

type TimeConstraint struct {
	start float64 // minutes precession
	end   float64 // minutes precession
}

func NewTimeConstraint(start float64, end float64) TimeConstraint {
	return TimeConstraint{
		start: start,
		end:   end,
	}
}

func (timeConstraint *TimeConstraint) isInRange(mintVal float64) bool {
	if mintVal >= timeConstraint.start && mintVal < timeConstraint.end {
		return true
	}
	return false
}

func (timeConstraint *TimeConstraint) isGreater(mintVal float64) bool {
	if mintVal > timeConstraint.end {
		return true
	}
	return false
}

// FlatHourly Model : hourly  price
type FlatHourly struct {
	price float64
}

func (flatHourly *FlatHourly) GetCost(parkingTime slot.ParkingTime) float64 {
	hours := parkingTime.CalculateHours()
	return math.Ceil(hours) * flatHourly.price
}

func NewFlatHourly(price float64) ModelCalculator {
	return &FlatHourly{price: price}
}

// DailyInterval Model : Daily  price
type DailyInterval struct {
	price float64
	TimeConstraint
}

func (dailyInterval *DailyInterval) GetCost(parkingTime slot.ParkingTime) float64 {
	mins := parkingTime.CalculateMinutes()
	if dailyInterval.isInRange(mins) {
		minutesOffSet := mins - dailyInterval.start
		return math.Ceil(MintoDays(minutesOffSet)) * dailyInterval.price
	}
	return 0
}

func NewDailyInterval(price float64, constraint TimeConstraint) ModelCalculator {
	return &DailyInterval{price: price, TimeConstraint: constraint}
}

// HourInterval Model :  fixed price for hour range
type HourInterval struct {
	TimeConstraint
	price float64
}

func (hourInterval *HourInterval) GetCost(parkingTime slot.ParkingTime) float64 {
	minutes := parkingTime.CalculateMinutes()
	if hourInterval.isInRange(minutes) {
		return hourInterval.price
	}
	return 0
}

func NewHourInterval(price float64, constraint TimeConstraint) ModelCalculator {
	return &HourInterval{price: price,
		TimeConstraint: constraint}
}

// InclusiveHourInterval Model :  include cost if park hour is greater than range value
type InclusiveHourInterval struct {
	TimeConstraint
	price float64
}

func (inclusiveHourInterval *InclusiveHourInterval) GetCost(parkingTime slot.ParkingTime) float64 {
	minutes := parkingTime.CalculateMinutes()
	if inclusiveHourInterval.isGreater(minutes) || inclusiveHourInterval.isInRange(minutes) {
		return inclusiveHourInterval.price
	}
	return 0
}

func NewInclusiveHourInterval(price float64, constraint TimeConstraint) ModelCalculator {
	return &InclusiveHourInterval{price: price,
		TimeConstraint: constraint}
}

// HourlyInterval Model :  hourly price for the range values
type HourlyInterval struct {
	TimeConstraint
	price float64
}

func (hourlyInterval *HourlyInterval) GetCost(parkingTime slot.ParkingTime) float64 {
	mins := parkingTime.CalculateMinutes()
	if hourlyInterval.isInRange(mins) {
		minutesOffSet := mins - hourlyInterval.start
		return math.Ceil(MintoHr(minutesOffSet)) * hourlyInterval.price
	}
	return 0
}

func NewHourlyInterval(price float64, constraint TimeConstraint) *HourlyInterval {
	return &HourlyInterval{price: price,
		TimeConstraint: constraint}
}

type Tariff interface {
	ModelCalculator
	Append(calculator ModelCalculator)
}

type BaseTariff struct {
	orderedTarrif []ModelCalculator
}

func (baseTariff *BaseTariff) Append(calculator ModelCalculator) {
	baseTariff.orderedTarrif = append(baseTariff.orderedTarrif, calculator)
}

// ParkingLotTarrif :  ordered list of tarrif Models one of model based on the
// definition
type ParkingLotTarrif struct {
	BaseTariff
}

func (parkingLotTarrif *ParkingLotTarrif) GetCost(parkingTime slot.ParkingTime) float64 {
	var cost float64
	for _, v := range parkingLotTarrif.orderedTarrif {
		cost = v.GetCost(parkingTime)
		if cost != 0 {
			break
		}
	}
	return cost
}

func NewParkingLotTarrif() Tariff {
	return &ParkingLotTarrif{}
}

// ParkingLotTarrifWithSum : ordered list of tarrif Models sums up all the model
// cost those in the range
type ParkingLotTarrifWithSum struct {
	BaseTariff
}

func (parkingLotTarrifWithSum *ParkingLotTarrifWithSum) GetCost(parkingTime slot.ParkingTime) float64 {
	var sum float64
	for _, v := range parkingLotTarrifWithSum.orderedTarrif {
		sum += v.GetCost(parkingTime)
	}
	return sum
}

func NewParkingLotTarrifWithSum() Tariff {
	return &ParkingLotTarrifWithSum{}
}
