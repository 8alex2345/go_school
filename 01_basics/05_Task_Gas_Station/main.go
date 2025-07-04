package main

import (
	"errors"
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
func (c *Car) RefuelCar(liter int) {
	if c.fuelLevel == c.capacity {
		fmt.Println("Бак полный заправка не нужна")
	} else {
		c.fuelLevel += liter
	}
}
func (g *GasStation) GetFuelGaz(liter int) error {
	if g.amountFuel < liter {
		return errors.New("заправка закрыта, нет бензина")
	}
	g.capacityGas -= liter
	fmt.Println("Топливо есть, можно заправляться")
	return nil
}

func CarRefuelingGas(station *GasStation, car *Car, liter int) {
	shortage := car.capacity - car.fuelLevel
	fmt.Println("До полного бака не хватает - ", shortage, "Введите количество которое хотите заправить - ", liter)
	if shortage < liter || liter < 1 {
		fmt.Println("Вы ввели недопустимое количество")
		return
	}
	err := station.GetFuelGaz(liter)
	if err != nil {
		fmt.Println("нет бензина на заправке")
		return
	}
	car.RefuelCar(liter)
	fmt.Println("Заправка закончена, залито в бак - ", liter)

}

func main() {
	gasStation := &GasStation{capacityGas: 100000, amountFuel: 50000}
	car := NewCar("Nissan", 100, 40)
	var liter int
	fmt.Scan(&liter)
	CarRefuelingGas(gasStation, car, liter)

}
