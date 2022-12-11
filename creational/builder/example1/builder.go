package main

type CarBuilder struct {
	car *Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{
		car: &Car{},
	}
}

func (c *CarBuilder) Make(make string) *CarBuilder {
	c.car.make = make
	return c
}

func (c *CarBuilder) Model(model string) *CarBuilder {
	c.car.model = model
	return c
}

func (c *CarBuilder) Year(year int) *CarBuilder {
	c.car.year = year
	return c
}

func (c *CarBuilder) Seats(seats int) *CarBuilder {
	c.car.seats = seats
	return c
}

func (c *CarBuilder) Build() *Car {
	return c.car
}
