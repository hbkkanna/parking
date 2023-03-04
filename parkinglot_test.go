package parking

import (
	"fmt"
	"github.com/hbkkanna/parking/slot"
	"github.com/hbkkanna/parking/tariff"
	"math"
	"testing"
	"time"
)

func TestParkingLotTest(t *testing.T) {
	// test case New Flat Hourly Model
	tariff1 := tariff.NewParkingLotTarrif()
	tariff1.Append(tariff.NewFlatHourly(10))

	tariff2 := tariff.NewParkingLotTarrif()
	tariff2.Append(tariff.NewFlatHourly(5))

	var configs []*ParkingConfig
	configs = append(configs, NewParkingConfig(slot.SCOOTER, 10, tariff1))
	configs = append(configs, NewParkingConfig(slot.SUV, 5, tariff2))
	configs = append(configs, NewParkingConfig(slot.SUV, 3, tariff2))

	parkingLot := NewParkingLot(configs)
	ticket, err := parkingLot.Park(slot.NewRoadVehicle(slot.SCOOTER))
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 1)
	receipt, _ := parkingLot.UnPark(ticket)
	if receipt.GetCost() != 10 {
		t.Errorf("wrong costing ")
	}
}

func TestMallParkingLot(t *testing.T) {

}

func TestStadiumParkingLot(t *testing.T) {

}

func TestAirportParkingLot(t *testing.T) {

}

/*
Mall
Per-hour flat fees
Vehicle Fee
Motorcycle 10
Car/SUV 20
Bus/Truck 50
*/
func getMallTariff() map[int]tariff.Tariff {
	tariffs := make(map[int]tariff.Tariff)

	scooterTariff := tariff.NewParkingLotTarrif()
	scooterTariff.Append(tariff.NewFlatHourly(10))
	tariffs[slot.SCOOTER] = scooterTariff

	carTariff := tariff.NewParkingLotTarrif()
	carTariff.Append(tariff.NewFlatHourly(20))
	tariffs[slot.SUV] = carTariff

	truckTariff := tariff.NewParkingLotTarrif()
	truckTariff.Append(tariff.NewFlatHourly(50))
	tariffs[slot.TRUCK] = truckTariff
	return tariffs
}

/*
Stadium
Flat rate up to a few hours and then per-hour rate. The total fee is the sum of all the previous
interval fees. No parking spots for buses/trucks at the stadium.
Vehicle Interval Fee
Motorcycle [0, 4) hours 30
[4, 12) hours 60
[12, Infinity) hours 100 per hour

Car/SUV [0, 4) hours 60
[4, 12) hours 120
[12, Infinity) hours 200 per hour
*/

func getStadiumTarrif() map[int]tariff.Tariff {
	tariffs := make(map[int]tariff.Tariff)

	scooterTariff := tariff.NewParkingLotTarrif()
	scooterTariff.Append(tariff.NewHourInterval(30, tariff.NewTimeConstraint(tariff.HrtoMinutes(0), tariff.HrtoMinutes(4))))
	scooterTariff.Append(tariff.NewHourInterval(60, tariff.NewTimeConstraint(tariff.HrtoMinutes(4), tariff.HrtoMinutes(12))))
	scooterTariff.Append(tariff.NewHourlyInterval(100, tariff.NewTimeConstraint(tariff.HrtoMinutes(12), math.MaxFloat64)))

	tariffs[slot.SCOOTER] = scooterTariff

	carTariff := tariff.NewParkingLotTarrif()
	carTariff.Append(tariff.NewHourInterval(60, tariff.NewTimeConstraint(tariff.HrtoMinutes(0), tariff.HrtoMinutes(4))))
	carTariff.Append(tariff.NewHourInterval(120, tariff.NewTimeConstraint(tariff.HrtoMinutes(4), tariff.HrtoMinutes(12))))
	carTariff.Append(tariff.NewHourlyInterval(200, tariff.NewTimeConstraint(tariff.HrtoMinutes(12), math.MaxFloat64)))
	tariffs[slot.SUV] = carTariff

	return tariffs
}
