package main

import "fmt"

type Car struct {
	make  string
	model string
	year  int
	seats int
}

func NewCar(make string, model string, year int, seats int) *Car {
	return &Car{
		make:  make,
		model: model,
		year:  year,
		seats: seats,
	}
}

func (car *Car) GetMake() string {
	return car.make
}

func (car *Car) GetModer() string {
	return car.model
}

func (car *Car) GetYear() int {
	return car.year
}

func (car *Car) GetSeats() int {
	return car.seats
}

func (car *Car) ToString() string {
	return fmt.Sprintf("make: %s\nmodel: %s\nyear: %d\nseats: %d\n",
		car.make, car.model, car.year, car.seats)
}
