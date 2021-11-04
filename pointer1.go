package main

import (
	"fmt"
)

func main() {

	x := 1

	y := &x

	*y = 2

	fmt.Println(x)
}
