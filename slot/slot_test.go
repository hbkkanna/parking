package slot

import (
	"fmt"
	"testing"
	"time"
)

func TestSlot(t *testing.T) {
	// reset and isFree
	slt := NewVehicleSlot(NewRoadVehicle(SCOOTER), 1)
	slt.Reset()
	if !slt.IsFree() {
		t.Errorf("isFree check failed ")
	}
	slt.SetInTime(time.Now())
	slt.SetInTime(time.Now().Add(time.Hour * 1))

	// clone test
	clonedSlot := CloneVehicleSlot(slt)
	// minimal check for reset
	if slt.GetInTime() != clonedSlot.GetInTime() || slt.GetOutTime() != clonedSlot.GetOutTime() {
		t.Errorf("clone failed")
	}
	if !clonedSlot.IsFree() {
		fmt.Println("clone working")
	}
}
