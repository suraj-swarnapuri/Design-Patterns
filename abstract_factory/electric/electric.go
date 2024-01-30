package electric

import inter "github.com/suraj-swarnapuri/Design-Patterns/abstract_factory/interfaces"

type Tesla struct {}

func (t Tesla) GetEngine() string {
	return "electric"
}

func (t Tesla) BeepHorn() string {
	return "beep beep"
}

type Bird struct {}

func (b *Bird) GetEngine() string {
	return "electric"
}

type ElectricFactory struct {}

func (e *ElectricFactory) CreateCar() inter.Car {
	return &Tesla{}
}

func (e *ElectricFactory) CreateBike() inter.Bike {
	return &Bird{}
}