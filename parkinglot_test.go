package parking

import (
	"fmt"
	"github.com/hbkkanna/parking/slot"
	"github.com/hbkkanna/parking/tariff"
	"math"
	"testing"
	"time"
)

// example 1

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

	scooterTariff := tariff.NewSingleTariffMatcher()
	scooterTariff.Append(tariff.NewEveryHour(10))
	tariffs[slot.SCOOTER] = scooterTariff

	carTariff := tariff.NewSingleTariffMatcher()
	carTariff.Append(tariff.NewEveryHour(20))
	tariffs[slot.SUV] = carTariff

	truckTariff := tariff.NewSingleTariffMatcher()
	truckTariff.Append(tariff.NewEveryHour(50))
	tariffs[slot.TRUCK] = truckTariff
	return tariffs
}

func SmallParkingLotConfig() []*ParkingConfig {
	var configs []*ParkingConfig
	configs = append(configs, NewParkingConfig(slot.SCOOTER, 2, getMallTariff()[slot.SCOOTER]))

	return configs
}

func TestSMallMallParkingLot(t *testing.T) {
	message := " ******** Mall parking lot case FAILED ******* "
	plot := NewParkingLot(SmallParkingLotConfig())
	lot, _ := plot.(*VehicleParkingLot)
	ticket, err := lot.park(slot.NewRoadVehicle(slot.SCOOTER), time.Now().Add(-time.Hour*3))
	if err == nil {
		receipt, _ := lot.UnPark(ticket)
		if receipt.GetCost() != 40 {
			t.Errorf(message)
		}
	}
	ticket, err = lot.park(slot.NewRoadVehicle(slot.SCOOTER), time.Now())
	if err == nil {
		receipt, _ := lot.UnPark(ticket)
		if receipt.GetCost() != 10 {
			t.Errorf(message)
		}
	}

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
func TestMallParkingLot(t *testing.T) {
	message := " ******** Mall parking lot case FAILED ******* "

	fmt.Println("----mall parking  case------ ")
	plot := NewParkingLot(MallParkingLotConfig())
	lot, _ := plot.(*VehicleParkingLot)
	ticket, err := lot.park(slot.NewRoadVehicle(slot.SCOOTER), time.Now().Add(-time.Hour*3))
	if err == nil {
		receipt, _ := lot.UnPark(ticket)
		if receipt.GetCost() != 40 {
			t.Errorf(message)
		}
	}
	ticket, err = lot.park(slot.NewRoadVehicle(slot.SUV), time.Now().Add(-time.Hour*2))
	if err == nil {
		receipt, _ := lot.UnPark(ticket)
		if receipt.GetCost() != 60 {
			t.Errorf(message)
		}
	}
	ticket, err = lot.park(slot.NewRoadVehicle(slot.TRUCK), time.Now())
	if err == nil {
		receipt, _ := lot.UnPark(ticket)
		if receipt.GetCost() != 50 {
			t.Errorf(message)
		}
	}
}

// example 3

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

	scooterTariff := tariff.NewMultipleTariffMatcher()
	scooterTariff.Append(tariff.NewPreviousHourInterval(30, tariff.NewTimeConstraint(tariff.HrtoMinutes(0), tariff.HrtoMinutes(4))))
	scooterTariff.Append(tariff.NewPreviousHourInterval(60, tariff.NewTimeConstraint(tariff.HrtoMinutes(4), tariff.HrtoMinutes(12))))
	scooterTariff.Append(tariff.NewEveryHourInInterval(100, tariff.NewTimeConstraint(tariff.HrtoMinutes(12), math.MaxFloat64)))

	tariffs[slot.SCOOTER] = scooterTariff

	carTariff := tariff.NewMultipleTariffMatcher()
	carTariff.Append(tariff.NewPreviousHourInterval(60, tariff.NewTimeConstraint(tariff.HrtoMinutes(0), tariff.HrtoMinutes(4))))
	carTariff.Append(tariff.NewPreviousHourInterval(120, tariff.NewTimeConstraint(tariff.HrtoMinutes(4), tariff.HrtoMinutes(12))))
	carTariff.Append(tariff.NewEveryHourInInterval(200, tariff.NewTimeConstraint(tariff.HrtoMinutes(12), math.MaxFloat64)))
	tariffs[slot.SUV] = carTariff

	return tariffs
}

func StadiumParkingLotConfig() []*ParkingConfig {
	var configs []*ParkingConfig
	stadiumTariff := getStadiumTarrif()
	configs = append(configs, NewParkingConfig(slot.SCOOTER, 1000, stadiumTariff[slot.SCOOTER]))
	configs = append(configs, NewParkingConfig(slot.SUV, 1500, stadiumTariff[slot.SUV]))
	return configs
}

func TestStadiumParkingLot(t *testing.T) {
	message := " ******** Stadium parking lot case FAILED ******* "

	plot := NewParkingLot(StadiumParkingLotConfig())
	lot, _ := plot.(*VehicleParkingLot)
	ticket, err := lot.park(slot.NewRoadVehicle(slot.SCOOTER), time.Now().Add(-time.Hour*13))
	if err == nil {
		receipt, _ := lot.UnPark(ticket)
		if receipt.GetCost() != 290 {
			t.Errorf(message)
		}
	}

	ticket, err = lot.park(slot.NewRoadVehicle(slot.SUV), time.Now().Add(-time.Hour*13))
	if err == nil {
		receipt, _ := lot.UnPark(ticket)
		if receipt.GetCost() != 580 {
			t.Errorf(message)
		}
	}

	// NO Space available - TRUCK not supported
	ticket, err = lot.park(slot.NewRoadVehicle(slot.TRUCK), time.Now().Add(-time.Hour*13))
	if err == nil {
		t.Errorf(message)
	}
}

//example 4

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

	scooterTariff := tariff.NewSingleTariffMatcher()
	scooterTariff.Append(tariff.NewHourInterval(0, tariff.NewTimeConstraint(tariff.HrtoMinutes(0), tariff.HrtoMinutes(1))))
	scooterTariff.Append(tariff.NewHourInterval(40, tariff.NewTimeConstraint(tariff.HrtoMinutes(1), tariff.HrtoMinutes(8))))
	scooterTariff.Append(tariff.NewHourInterval(60, tariff.NewTimeConstraint(tariff.HrtoMinutes(8), tariff.HrtoMinutes(24))))
	scooterTariff.Append(tariff.NewEveryDay(80, tariff.NewTimeConstraint(tariff.DaytoMinutes(0), math.MaxFloat64)))

	tariffs[slot.SCOOTER] = scooterTariff

	carTariff := tariff.NewSingleTariffMatcher()
	carTariff.Append(tariff.NewHourInterval(60, tariff.NewTimeConstraint(tariff.HrtoMinutes(0), tariff.HrtoMinutes(12))))
	carTariff.Append(tariff.NewHourInterval(80, tariff.NewTimeConstraint(tariff.HrtoMinutes(12), tariff.HrtoMinutes(24))))
	carTariff.Append(tariff.NewEveryDay(100, tariff.NewTimeConstraint(tariff.DaytoMinutes(0), math.MaxFloat64)))
	tariffs[slot.SUV] = carTariff

	return tariffs
}

func AirportParkingLotConfig() []*ParkingConfig {
	var configs []*ParkingConfig
	airportTariff := getAirportTarrif()
	configs = append(configs, NewParkingConfig(slot.SCOOTER, 200, airportTariff[slot.SCOOTER]))
	configs = append(configs, NewParkingConfig(slot.SUV, 500, airportTariff[slot.SUV]))
	return configs
}

func TestAirportParkingLot(t *testing.T) {
	message := " ******** Airport parking lot case FAILED ******* "
	plot := NewParkingLot(AirportParkingLotConfig())
	lot, _ := plot.(*VehicleParkingLot)
	ticket, err := lot.park(slot.NewRoadVehicle(slot.SCOOTER), time.Now())
	if err == nil {
		receipt, _ := lot.UnPark(ticket)
		if receipt.GetCost() != 0 {
			t.Errorf(message)
		}
	}

	ticket, err = lot.park(slot.NewRoadVehicle(slot.SCOOTER), time.Now().Add(-time.Hour*24))
	if err == nil {
		receipt, _ := lot.UnPark(ticket)
		if receipt.GetCost() != 160 {
			t.Errorf(message)
		}
	}
}

func TestParkingLotFullCase(t *testing.T) {
	plot := NewParkingLot(SmallParkingLotConfig())
	lot, _ := plot.(*VehicleParkingLot)
	fmt.Println("-----parking lot full case------ ")
	lot.park(slot.NewRoadVehicle(slot.SCOOTER), time.Now())
	lot.park(slot.NewRoadVehicle(slot.SCOOTER), time.Now())
	_, err := lot.park(slot.NewRoadVehicle(slot.SCOOTER), time.Now())
	if err == nil {
		t.Errorf("Parking lot full case failed ")
	}

}
