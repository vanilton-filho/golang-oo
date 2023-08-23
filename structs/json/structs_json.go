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

	car := Car{"Gol", 2017, "Yellow"}

	result, _ := json.Marshal(car)
	fmt.Println(string(result))

}
