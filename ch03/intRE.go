package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	s := arguments[1]
	ret := matchInt(s)
	fmt.Println(ret)

	integers := 0
	nonintegers := 0
	for _, arg := range os.Args[1:] {
		s := arg
		ret := matchInt(s)
		fmt.Println(ret)
		if ret {
			integers++
		} else {
			nonintegers++
		}
	}
	fmt.Println("Totals ", integers, " ", nonintegers)
}
