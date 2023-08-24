package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string `json:"name"`
	Year  int    `json:"-"`
	Color string
}

func main() {

	var car Car
	data := []byte(`{"name": "Gol", "Color": "Black"}`)

	json.Unmarshal(data, &car)
	fmt.Println(car.Name, car.Year, car.Color)

}
