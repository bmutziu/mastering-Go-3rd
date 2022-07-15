package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	s := arguments[1]
	ret := matchNameSur(s)
	fmt.Println(ret)

	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
		s := arg
		ret := matchNameSur(s)
		fmt.Println(ret)
	}
}
