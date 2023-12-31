package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

type SuperCar struct {
	Car
	CanFly bool
	Name   string
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s\nYear: %d\nColor: %s\n", c.Name, c.Year, c.Color)
}

func main() {

	// car := Car{"Corolla", 2017, "White"}
	// car2 := Car{"BMW 320i", 2017, "Black"}
	sCar := SuperCar{Car: Car{"Fusca", 2030, "Blue"}, CanFly: true, Name: "Elantra"}

	// fmt.Println(car.Name, car.Year, car.Color)
	// fmt.Println(car2.Name, car2.Year, car.Color)
	// fmt.Println(car.info())

	fmt.Println(sCar.Name)
	fmt.Println(sCar.info())
}
