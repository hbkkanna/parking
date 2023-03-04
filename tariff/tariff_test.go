package tariff

import (
	"fmt"
	"github.com/hbkkanna/parking/slot"
	"math"
	"testing"
	"time"
)

func TestFlatHourly(t *testing.T) {

	// test case  Flat Hourly Model
	tariff1 := NewParkingLotTarrif()
	tariff1.Append(NewFlatHourly(10))
	newTicket := slot.NewTicket(1, slot.NewVehicleSlot(slot.Vehicles[slot.SCOOTER], 2))

	cur := time.Now()
	newTicket.SetInTime(cur)
	newTicket.SetOutTime(cur.Add(time.Hour*3 + time.Second*10))
	if tariff1.GetCost(newTicket) != 40 {
		t.Errorf("flat hourly failed ")
	} else {
		fmt.Println(tariff1.GetCost(newTicket))
	}

	// 0 difference case
	cur = time.Now()
	newTicket.SetInTime(cur)
	newTicket.SetOutTime(time.Now())
	fmt.Printf("%.2f", newTicket.CalculateMinutes())
	if tariff1.GetCost(newTicket) != 10 {
		t.Errorf("flat hourly failed ")
	} else {
		fmt.Println(tariff1.GetCost(newTicket))
	}
}

func TestHourInterval(t *testing.T) {
	// test case HourlyInterval Model
	tariff1 := NewParkingLotTarrif()
	tariff1.Append(NewHourInterval(10, TimeConstraint{start: HrtoMinutes(0), end: HrtoMinutes(4)}))
	tariff1.Append(NewHourInterval(20, TimeConstraint{start: HrtoMinutes(4), end: HrtoMinutes(8)}))
	tariff1.Append(NewHourInterval(30, TimeConstraint{start: HrtoMinutes(8), end: HrtoMinutes(12)}))

	newTicket := slot.NewTicket(1, slot.NewVehicleSlot(slot.Vehicles[slot.SCOOTER], 2))
	newTicket.SetInTime(time.Now())
	newTicket.SetOutTime(time.Now().Add(time.Hour * 5))

	if tariff1.GetCost(newTicket) != 20 {
		fmt.Println(tariff1.GetCost(newTicket))
		t.Errorf("Hour Interval Model failed ")
	} else {
		fmt.Println(tariff1.GetCost(newTicket))
	}

	// minutes test case
	newTicket = slot.NewTicket(1, slot.NewVehicleSlot(slot.Vehicles[slot.SCOOTER], 2))
	newTicket.SetInTime(time.Now())
	newTicket.SetOutTime(time.Now().Add(time.Hour * 11).Add(time.Minute * 59))

	if tariff1.GetCost(newTicket) != 30 {
		fmt.Println(newTicket.CalculateHours())
		t.Errorf("Hour Interval Model failed ")
	} else {
		fmt.Println(tariff1.GetCost(newTicket))
	}

	// not in range case
	newTicket = slot.NewTicket(1, slot.NewVehicleSlot(slot.Vehicles[slot.SCOOTER], 2))
	val := time.Now()
	newTicket.SetInTime(val)
	newTicket.SetOutTime(val.Add(time.Hour * 12))
	if tariff1.GetCost(newTicket) != 0 {
		fmt.Printf("%.2f", newTicket.CalculateHours())
		t.Errorf("Hour Interval Model failed ")
	} else {
		fmt.Println(tariff1.GetCost(newTicket))
	}

}

func TestDailyInterval(t *testing.T) {
	// test case DailyInterval Model
	tariff1 := NewParkingLotTarrif()
	tariff1.Append(NewHourInterval(10, TimeConstraint{start: HrtoMinutes(0), end: HrtoMinutes(8)}))
	tariff1.Append(NewHourInterval(20, TimeConstraint{start: HrtoMinutes(8), end: HrtoMinutes(24)}))

	tariff1.Append(NewDailyInterval(30, TimeConstraint{start: DaytoMinutes(0), end: math.MaxFloat64}))

	newTicket := slot.NewTicket(1, slot.NewVehicleSlot(slot.Vehicles[slot.SCOOTER], 2))
	newTicket.SetInTime(time.Now())
	newTicket.SetOutTime(time.Now().Add(time.Hour * 24 * 5)) // 6 days

	// 6 * 30
	if tariff1.GetCost(newTicket) != 180 {
		//fmt.Println(math.Ceil(MintoDays(newTicket.CalculateMinutes())))
		//fmt.Println(newTicket.CalculateHours())
		t.Errorf("Hourly Interval Model failed ")
	}
	fmt.Println(tariff1.GetCost(newTicket))

}

func TestHourlyInterval(t *testing.T) {
	// test case HourlyInterval Model
	tariff1 := NewParkingLotTarrif()
	tariff1.Append(NewHourInterval(10, TimeConstraint{start: HrtoMinutes(0), end: HrtoMinutes(8)}))
	tariff1.Append(NewHourInterval(20, TimeConstraint{start: HrtoMinutes(8), end: HrtoMinutes(12)}))
	tariff1.Append(NewHourlyInterval(30, TimeConstraint{start: HrtoMinutes(12), end: math.MaxFloat64}))

	newTicket := slot.NewTicket(1, slot.NewVehicleSlot(slot.Vehicles[slot.SCOOTER], 2))
	newTicket.SetInTime(time.Now())
	newTicket.SetOutTime(time.Now().Add(time.Hour * 15))
	fmt.Println(tariff1.GetCost(newTicket))

	// (16-12) * 30
	if tariff1.GetCost(newTicket) != 120 {
		fmt.Println(newTicket.CalculateHours())
		t.Errorf("Hourly Interval Model failed ")
	} else {
		fmt.Println(tariff1.GetCost(newTicket))
	}
}

func TestParkingLotTariffSum(t *testing.T) {
	// test parking Lot Sum & inclusive model
	tariff1 := NewParkingLotTarrifWithSum()
	tariff1.Append(NewInclusiveHourInterval(10, TimeConstraint{start: HrtoMinutes(0), end: HrtoMinutes(8)}))
	tariff1.Append(NewInclusiveHourInterval(20, TimeConstraint{start: HrtoMinutes(8), end: HrtoMinutes(12)}))
	tariff1.Append(NewHourlyInterval(30, TimeConstraint{start: HrtoMinutes(12), end: math.MaxFloat64}))

	newTicket := slot.NewTicket(1, slot.NewVehicleSlot(slot.Vehicles[slot.SCOOTER], 2))
	newTicket.SetInTime(time.Now())
	newTicket.SetOutTime(time.Now().Add(time.Hour * 15))
	fmt.Println(tariff1.GetCost(newTicket))

	// (16-12) * 30 + 10 + 20 = 150
	if tariff1.GetCost(newTicket) != 150 {
		t.Errorf("Parking lot tariff failed ")
	}
	fmt.Println(tariff1.GetCost(newTicket))

}
