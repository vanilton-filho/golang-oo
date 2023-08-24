package main

import "fmt"

type vehicle interface {
	start() string
}

type car struct {
	name string
}

func (c car) start() string {
	return "The car " + c.name + " has been started."
}

type motorcycle struct {
	name string
}

func (mc motorcycle) start() string {
	return "The motorcycle " + mc.name + " has been started."
}

func startVehicle(v vehicle) string {
	return v.start()
}

func main() {
	c := car{"Gol"}
	m := motorcycle{"Broz"}

	fmt.Println(startVehicle(c))
	fmt.Println(startVehicle(m))

}
