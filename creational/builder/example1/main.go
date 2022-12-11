package main

func main() {
	car := NewCarBuilder().
		Make("Mercedes-Benz").
		Model("W212").
		Seats(4).
		Year(2012).
		Build()

	print(car.ToString())
}
