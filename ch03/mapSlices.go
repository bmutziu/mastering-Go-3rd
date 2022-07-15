package main

import "fmt"

func main() {
	// map to slices

	aMap := make(map[int]string)
	aMap[0] = "Zero"
	aMap[1] = "Unu"
	aMap[2] = "Doi"

	keys := make([]int, 0, len(aMap))
	values := make([]string, 0, len(aMap))

	for k, v := range aMap {
		fmt.Println("key:", k, "value:", v)
		keys = append(keys, k)
		values = append(values, v)
	}

	fmt.Println(keys, "with length", len(keys))
	fmt.Println(values, "with length", len(values))
}
