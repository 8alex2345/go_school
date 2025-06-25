package main

import (
	"fmt"
)

type Car struct {
	name      string
	capacity  int
	fuelLevel int
}
type GasStation struct {
	amountFuel  int
	capacityGas int
}

func NewCar(name string, capacity, fuelLevel int) *Car {
	if fuelLevel < 0 || capacity < 0 {
		fmt.Println("Введены отрицательные числа")
	}
	return &Car{
		name:      name,
		capacity:  capacity,
		fuelLevel: fuelLevel,
	}
}

type GasSstationImpl interface {
	RefuelCar(shortage int, liter int)
	GetFuelGaz(gas *GasStation) int
}

func (c *Car) RefuelCar(shortage, liter int) {
	shortage = c.capacity - c.fuelLevel
	fmt.Println("До полного бака не хватает - ", shortage, "Введите количество которое хотите заправить - ", liter)
	if shortage < liter || liter < 1 {
		fmt.Println("Вы ввели недопустимое количество")
	}
	c.fuelLevel += liter
	fmt.Println("Заправка закончена, залито в бак - ", liter)

}
func (g *GasStation) GetFuelGaz(liter int) bool {
	if g.amountFuel < liter {
		fmt.Println("Заправка закрыта, нет бензина")
		return false
	}
	g.capacityGas -= liter
	fmt.Println("Топливо есть, можно заправляться")
	return true
}

func Refueling(station *GasStation, car *Car, liter int) {
	if car.fuelLevel < liter {

	}
}

func main() {
	gasStation := &GasStation{capacityGas: 100000, amountFuel: 50000}
	car := NewCar("Nissan", 100, 40)
	var liter int
	fmt.Scan(&liter)
	Refueling(gasStation, car, liter)

}
