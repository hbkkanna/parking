package slot

import (
	"fmt"
	"testing"
	"time"
)

func TestParkingTime(t *testing.T) {
	var parkingTime ParkingTime
	parkingTime = &VehicleParkingTime{}

	// out time error check
	timeVal := time.Now().Add(time.Minute * 2)
	parkingTime.SetInTime(timeVal)
	timeVal = time.Now()
	err := parkingTime.SetOutTime(timeVal)
	if err == nil {
		t.Errorf("Failed to validate  ")
	}

	// calculate hour and minute
	timeVal = time.Now()
	parkingTime.SetInTime(timeVal)
	parkingTime.SetOutTime(timeVal.Add(time.Hour * 2))

	if hr := parkingTime.CalculateHours(); hr != 2 {
		t.Errorf("Failed : to calcuate  hours %f ", hr)
	} else {
		fmt.Printf(" CalculateHours output  %f \n", hr)
	}

	if min := parkingTime.CalculateMinutes(); min != 120 {
		t.Errorf("Failed : to calcuate  minutes ")
	} else {
		fmt.Printf(" CalculateDays output  %f \n", min)
	}

}
