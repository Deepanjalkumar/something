package main

import (
	"fmt"
)

func main() {

	x := "name"

	var x1 *string = &x

	fmt.Println(*x1)

}
