package main

import "fmt"

func main() {
	slice := []byte{0, 1, 2, 3}
	array := [8]byte{slice[0], slice[1], slice[2], slice[3]}
	fmt.Println(array)

	slice2 := []byte{4, 5, 6, 7}
	for index, bite := range slice2 {
		array[index+len(slice)] = bite
	}
	fmt.Println(array)

	slice = append(slice, slice2...)

	copy(array[:], slice) // NOTE the slicing of the array
	fmt.Println(slice)
	fmt.Println(array)
}
