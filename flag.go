package main

import (
	"flag"
	"fmt"
)

func Sum(x int, y int) (int, error) {
	z := x + y
	return z, nil
}

func main() {

	first := flag.Int("first", 0, "enter first number")
	second := flag.Int("second", 0, "enter second number")
	flag.Parse()
	value, _ := Sum(*first, *second)
	fmt.Println(value)
}
