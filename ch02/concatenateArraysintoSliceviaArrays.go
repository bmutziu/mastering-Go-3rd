package main

import (
	"fmt"
)

func main() {
	var seven [7]int

	five := [5]int{1, 2, 3, 4, 5}
	two := [2]int{6, 7}

	s := seven[:0]
	fmt.Println("s: ", s)
	s = append(s, five[:]...)
	s = append(s, two[:]...)
	fmt.Println("seven:", seven)
	fmt.Println("s:", s)

	var sevens [7]int

	copy(sevens[:], five[:])
	copy(sevens[len(five):], two[:])
	fmt.Println("sevens: ", sevens)
}
