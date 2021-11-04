package main

import (
	"fmt"
)

func main() {

	array := []int{1, 2, 3, 4, 5}

	for index, _ := range array {
		fmt.Println(array[index])
	}
}
