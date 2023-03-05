package parking

import (
	"fmt"
	"github.com/hbkkanna/parking/slot"
	"github.com/hbkkanna/parking/tariff"
	"math"
	"testing"
	"time"
)

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

	scooterTariff := tariff.NewParkingLotTarrifWithSum()
	scooterTariff.Append(tariff.NewInclusiveHourInterval(30, tariff.NewTimeConstraint(tariff.HrtoMinutes(0), tariff.HrtoMinutes(4))))
	scooterTariff.Append(tariff.NewInclusiveHourInterval(60, tariff.NewTimeConstraint(tariff.HrtoMinutes(4), tariff.HrtoMinutes(12))))
	scooterTariff.Append(tariff.NewHourlyInterval(100, tariff.NewTimeConstraint(tariff.HrtoMinutes(12), math.MaxFloat64)))

	tariffs[slot.SCOOTER] = scooterTariff

	carTariff := tariff.NewParkingLotTarrifWithSum()
	carTariff.Append(tariff.NewInclusiveHourInterval(60, tariff.NewTimeConstraint(tariff.HrtoMinutes(0), tariff.HrtoMinutes(4))))
	carTariff.Append(tariff.NewInclusiveHourInterval(120, tariff.NewTimeConstraint(tariff.HrtoMinutes(4), tariff.HrtoMinutes(12))))
	carTariff.Append(tariff.NewHourlyInterval(200, tariff.NewTimeConstraint(tariff.HrtoMinutes(12), math.MaxFloat64)))
	tariffs[slot.SUV] = carTariff

	return tariffs
}

/*
Airport
Flat rate up to one day. Then per-day rate. There is no summing up of the previous interval
fees. No parking spots for buses/trucks at the airport.
Vehicle Interval Fee
Motorcycle [0, 1) hours Free
[1, 8) hours 40
[8, 24) hours 60
Each day 80
Car/SUV [0, 12) hours 60
[12, 24) hours 80
Each day 100
*/
func getAirportTarrif() map[int]tariff.Tariff {
	tariffs := make(map[int]tariff.Tariff)

	scooterTariff := tariff.NewParkingLotTarrif()
	scooterTariff.Append(tariff.NewHourInterval(0, tariff.NewTimeConstraint(tariff.HrtoMinutes(0), tariff.HrtoMinutes(1))))
	scooterTariff.Append(tariff.NewHourInterval(40, tariff.NewTimeConstraint(tariff.HrtoMinutes(1), tariff.HrtoMinutes(8))))
	scooterTariff.Append(tariff.NewHourInterval(60, tariff.NewTimeConstraint(tariff.HrtoMinutes(8), tariff.HrtoMinutes(24))))
	scooterTariff.Append(tariff.NewDailyInterval(80, tariff.NewTimeConstraint(tariff.DaytoMinutes(0), math.MaxFloat64)))

	tariffs[slot.SCOOTER] = scooterTariff

	carTariff := tariff.NewParkingLotTarrif()
	carTariff.Append(tariff.NewHourInterval(60, tariff.NewTimeConstraint(tariff.HrtoMinutes(0), tariff.HrtoMinutes(12))))
	carTariff.Append(tariff.NewHourInterval(80, tariff.NewTimeConstraint(tariff.HrtoMinutes(12), tariff.HrtoMinutes(24))))
	carTariff.Append(tariff.NewDailyInterval(100, tariff.NewTimeConstraint(tariff.DaytoMinutes(0), math.MaxFloat64)))
	tariffs[slot.SUV] = carTariff

	return tariffs
}

// example 1
func SmallParkingLotConfig() []*ParkingConfig {
	var configs []*ParkingConfig
	configs = append(configs, NewParkingConfig(slot.SCOOTER, 2, getMallTariff()[slot.SCOOTER]))

	return configs
}

// example 2
func MallParkingLotConfig() []*ParkingConfig {
	var configs []*ParkingConfig

	mallTariff := getMallTariff()
	configs = append(configs, NewParkingConfig(slot.SCOOTER, 100, mallTariff[slot.SCOOTER]))
	configs = append(configs, NewParkingConfig(slot.SUV, 80, mallTariff[slot.SUV]))
	configs = append(configs, NewParkingConfig(slot.TRUCK, 100, mallTariff[slot.TRUCK]))

	return configs
}

// example 3
func StadiumParkingLotConfig() []*ParkingConfig {
	var configs []*ParkingConfig
	stadiumTariff := getStadiumTarrif()
	configs = append(configs, NewParkingConfig(slot.SCOOTER, 1000, stadiumTariff[slot.SCOOTER]))
	configs = append(configs, NewParkingConfig(slot.SUV, 1500, stadiumTariff[slot.SUV]))
	return configs
}

// example 3
func AirportParkingLotConfig() []*ParkingConfig {
	var configs []*ParkingConfig
	airportTariff := getAirportTarrif()
	configs = append(configs, NewParkingConfig(slot.SCOOTER, 200, airportTariff[slot.SCOOTER]))
	configs = append(configs, NewParkingConfig(slot.SUV, 500, airportTariff[slot.SUV]))
	return configs
}

func TestParkingLot(t *testing.T) {
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
	plot := NewParkingLot(SmallParkingLotConfig())
	lot, _ := plot.(*VehicleParkingLot)
	ticket, err := lot.park(slot.NewRoadVehicle(slot.SCOOTER), time.Now().Add(-time.Hour*3))
	if err == nil {
		lot.UnPark(ticket)
	}
	ticket, err = lot.Park(slot.NewRoadVehicle(slot.SCOOTER))
	if err == nil {
		lot.UnPark(ticket)
	}
	ticket, err = lot.Park(slot.NewRoadVehicle(slot.SCOOTER))
	if err == nil {
		lot.UnPark(ticket)
	}

	lot.Park(slot.NewRoadVehicle(slot.SCOOTER))
	lot.Park(slot.NewRoadVehicle(slot.SCOOTER))
	lot.Park(slot.NewRoadVehicle(slot.SCOOTER))

}

func TestStadiumParkingLot(t *testing.T) {

}

func TestAirportParkingLot(t *testing.T) {

}
