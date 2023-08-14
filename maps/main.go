package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
	fmt.Println(m)

	delete(m, "c")
	fmt.Println(m["c"])
	fmt.Println(m)

	_, exists := m["c"]
	fmt.Println(exists)
	value, exists := m["b"]
	fmt.Println(value, exists)

	var x = map[string]int{
		"golang": 10,
	}
	fmt.Println(x)

	newMap := map[string]int{"x": 5}
	fmt.Println(newMap)

	if v, exists := m["a"]; exists {
		fmt.Println("Golang is", v)
	} else {
		fmt.Println(":(")
	}
}
