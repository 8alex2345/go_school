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
	RefuelCar(liter int)
	GetFuelGaz(gas *GasStation) int
}

func (c *Car) RefuelCar(liter int) {
	if c.fuelLevel+liter > c.capacity {
		c.fuelLevel = c.capacity
		fmt.Println("Бак заполнен")
	} else {
		c.fuelLevel += liter
	}
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

func CarRefuelingGas(station *GasStation, car *Car, liter int) {
	shortage := car.capacity - car.fuelLevel
	fmt.Println("До полного бака не хватает - ", shortage, "Введите количество которое хотите заправить - ", liter)
	if shortage < liter || liter < 1 {
		fmt.Println("Вы ввели недопустимое количество")
	}
	if station.GetFuelGaz(liter) {
		car.RefuelCar(liter)
		fmt.Println("Заправка закончена, залито в бак - ", liter)
	} else {
		fmt.Println("На заправке недостоточно топлива")
	}

}

func main() {
	gasStation := &GasStation{capacityGas: 100000, amountFuel: 50000}
	car := NewCar("Nissan", 100, 40)
	var liter int
	fmt.Scan(&liter)
	CarRefuelingGas(gasStation, car, liter)

}
