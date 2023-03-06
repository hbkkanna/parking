package tariff

import (
	"github.com/hbkkanna/parking/slot"
	"math"
)

type ModelCalculator interface {
	GetCost(parkingTime slot.ParkingTime) float64
}

type TimeConstraint struct {
	start float64 // nanoseconds precession in minutes
	end   float64
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

// EveryHour Model : hourly  price
type EveryHour struct {
	price float64
}

func (everyHour *EveryHour) GetCost(parkingTime slot.ParkingTime) float64 {
	hours := parkingTime.CalculateHours()
	return math.Ceil(hours) * everyHour.price
}

func NewEveryHour(price float64) ModelCalculator {
	return &EveryHour{price: price}
}

// EveryDay Model : Daily  price
type EveryDay struct {
	price float64
	TimeConstraint
}

func (everyDay *EveryDay) GetCost(parkingTime slot.ParkingTime) float64 {
	mins := parkingTime.CalculateMinutes()
	if everyDay.isInRange(mins) {
		minutesOffSet := mins - everyDay.start
		return math.Ceil(MintoDays(minutesOffSet)) * everyDay.price
	}
	return -1
}

func NewEveryDay(price float64, constraint TimeConstraint) ModelCalculator {
	return &EveryDay{price: price, TimeConstraint: constraint}
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
	return -1
}

func NewHourInterval(price float64, constraint TimeConstraint) ModelCalculator {
	return &HourInterval{price: price,
		TimeConstraint: constraint}
}

// PreviousHourInterval Model :  include cost if park hour is greater than range value
type PreviousHourInterval struct {
	TimeConstraint
	price float64
}

func (previousHourInterval *PreviousHourInterval) GetCost(parkingTime slot.ParkingTime) float64 {
	minutes := parkingTime.CalculateMinutes()
	if previousHourInterval.isGreater(minutes) || previousHourInterval.isInRange(minutes) {
		return previousHourInterval.price
	}
	return -1
}

func NewPreviousHourInterval(price float64, constraint TimeConstraint) ModelCalculator {
	return &PreviousHourInterval{price: price,
		TimeConstraint: constraint}
}

// EveryHourInInterval Model :  hourly price for the range values
type EveryHourInInterval struct {
	TimeConstraint
	price float64
}

func (everyHourInInterval *EveryHourInInterval) GetCost(parkingTime slot.ParkingTime) float64 {
	mins := parkingTime.CalculateMinutes()
	if everyHourInInterval.isInRange(mins) {
		minutesOffSet := mins - everyHourInInterval.start
		return math.Ceil(MintoHr(minutesOffSet)) * everyHourInInterval.price
	}
	return -1
}

func NewEveryHourInInterval(price float64, constraint TimeConstraint) *EveryHourInInterval {
	return &EveryHourInInterval{price: price,
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

// SingleTariffMatcher : matches with single model in ordered list
type SingleTariffMatcher struct {
	BaseTariff
}

func (singleTariffMatcher *SingleTariffMatcher) GetCost(parkingTime slot.ParkingTime) float64 {
	var cost float64
	for _, v := range singleTariffMatcher.orderedTarrif {
		cost = v.GetCost(parkingTime)
		if cost > -1 {
			break
		}
	}
	return cost
}

func NewSingleTariffMatcher() Tariff {
	return &SingleTariffMatcher{}
}

// MultipleTariffMatcher : sums up all the matching models
type MultipleTariffMatcher struct {
	BaseTariff
}

func (multipleTariffMatcher *MultipleTariffMatcher) GetCost(parkingTime slot.ParkingTime) float64 {
	var sum float64
	for _, v := range multipleTariffMatcher.orderedTarrif {
		sum += v.GetCost(parkingTime)
	}
	return sum
}

func NewMultipleTariffMatcher() Tariff {
	return &MultipleTariffMatcher{}
}
