package main

import "fmt"

type Names []interface{}

func (n *Names) load() {
	fmt.Println(n)
	*n = Names{
		"Wesley",
		"Sarah",
		"Davi",
		10,
	}
	fmt.Println(n)
}

func (n Names) printNames() {
	fmt.Println(n)
}

func main() {

	var names Names

	names.load()
	names.printNames()
}
