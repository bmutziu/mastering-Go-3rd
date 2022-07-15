package main

import "fmt"

func main() {
	var elements = [4]string{"abc", "def", "fgi", "adi"}
	elementMap := make(map[string]string)

	for i := 0; i < len(elements); i += 2 {
		elementMap[elements[i]] = elements[i+1]
	}

	for key, v := range elementMap {
		fmt.Println("key:", key, "value:", v)
	}
}
