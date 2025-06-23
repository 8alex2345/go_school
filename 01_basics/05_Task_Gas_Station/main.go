package main

import "fmt"

type Car struct {
	Name      string
	Capacity  int
	FuelLevel int
}

type GasStation struct {
	refuel int
	status string
}
type GasSstationImpl interface {
	refuelCar(car *Car) string
}

func (g *GasStation) refuelCar(car *Car) string {
	if car.Capacity > car.FuelLevel {
		g.refuel = car.Capacity - car.FuelLevel
		car.FuelLevel += g.refuel
		g.status = "Идет заправка"
		return g.status
	}

	return "Машина заправлена"

}
func main() {
	car1 := &Car{Name: "Nissan", Capacity: 100, FuelLevel: 40}
	gasStation := &GasStation{}
	status := gasStation.refuelCar(car1)
	fmt.Println(status)

	fmt.Println("Уровеь топлива - ", car1.FuelLevel)

	status = gasStation.refuelCar(car1)
	fmt.Println(status)

}
