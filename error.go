package main

import (
	"errors"
	"fmt"
)

func Sum(x int, y int) (int, error) {

	if x <= 0 && y <= 0 {
		return 0, errors.New("Do not pass value less than or equal to 0")
	} else {
		z := x + y
		return z, nil

	}

}

func main() {

	value, err := Sum(0, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

}
