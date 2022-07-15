package main

import (
	"fmt"
	"os"
)

type recordbis struct {
	Index int
	Value string
}

func main() {
	Sbis := []recordbis{}
	fmt.Println("len", len(os.Args))

	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
		temp := recordbis{Index: i, Value: arg}
		Sbis = append(Sbis, temp)
	}

	fmt.Println(Sbis, "with length", len(Sbis))
}
