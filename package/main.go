package main

import (
	"fmt"
	"golang-oo/package/car"
)

func main() {

	car := car.Car{Name: "Fusca", Color: "Black"}
	fmt.Println(car.Start())
}
