package main

import "fmt"

func main() {

	slice := make([]int, 2, 5)
	// slice[3] = 2

	// slice = append(slice, 1, 10, 42, 23, 21, 33, 2, 4, 55, 33, 0)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// for i := 0; i < 20; i++ {
	// 	slice = append(slice, i)
	// 	fmt.Println("Len: ", len(slice), "Cap: ", cap(slice))
	// }

	// slice = append(slice, 1, 3, 4, 5, 6, 7, 8, 9, 10)
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	sliceString := []string{
		"Hello",
		"World",
		"Much",
		"Better",
	}

	fmt.Println(sliceString)
	fmt.Println(sliceString[0])
	fmt.Println(sliceString[:2])
	fmt.Println(sliceString[1:2])
	fmt.Println(sliceString[2:4])
	fmt.Println(sliceString[2:])
}
