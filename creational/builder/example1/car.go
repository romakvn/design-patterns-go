package main

import "fmt"

type Car struct {
	make  string
	model string
	year  int
	seats int
}

func (car *Car) ToString() string {
	return fmt.Sprintf("make: %s\nmodel: %s\nyear: %d\nseats: %d\n",
		car.make, car.model, car.year, car.seats)
}
