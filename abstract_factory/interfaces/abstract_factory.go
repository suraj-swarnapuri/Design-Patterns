package interfaces


type Car interface {
	GetEngine() string
	BeepHorn() string
}

type Bike interface {
	GetEngine() string
}

type AbstractFactory interface {
	CreateCar() Car
	CreateBike() Bike
}