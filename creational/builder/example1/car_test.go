package main

import "testing"

func InitializeCarObject() *Car {
	return NewCar(
		"Subaru",
		"Forester",
		2015,
		4,
	)
}

func TestGetMakeMethod(t *testing.T) {
	car := InitializeCarObject()

	if car.GetMake() != "Subaru" {
		t.Errorf("Expected result: Subaru, got %s", car.GetMake())
	}
}

func TestGetYearMethod(t *testing.T) {
	car := InitializeCarObject()

	if car.GetYear() != 2015 {
		t.Errorf("Expected result: Subaru, got %d", car.GetYear())
	}
}
