package gas

import inter "github.com/suraj-swarnapuri/Design-Patterns/abstract_factory/interfaces"

type BMW struct {}

func (t BMW) GetEngine() string {
	return "gas"
}

func (t BMW) BeepHorn() string {
	return "honk"
}

type Razor struct {}

func (b *Razor) GetEngine() string {
	return "human"
}

type GasFactory struct {}

func (e *GasFactory) CreateCar() inter.Car {
	return &BMW{}
}

func (e *GasFactory) CreateBike() inter.Bike {
	return &Razor{}
}