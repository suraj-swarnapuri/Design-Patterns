package main

import (
	electric "github.com/suraj-swarnapuri/Design-Patterns/abstract_factory/electric"
	gas "github.com/suraj-swarnapuri/Design-Patterns/abstract_factory/gas"
	inter "github.com/suraj-swarnapuri/Design-Patterns/abstract_factory/interfaces"
)





func main() {
	var factory inter.AbstractFactory

	factory = &electric.ElectricFactory{}
	car := factory.CreateCar()
	bike := factory.CreateBike()

	println(car.GetEngine())
	println(car.BeepHorn())

	println(bike.GetEngine())

	factory = &gas.GasFactory{}
	car = factory.CreateCar()
	bike = factory.CreateBike()

	println(car.GetEngine())
	println(car.BeepHorn())

	println(bike.GetEngine())
}

