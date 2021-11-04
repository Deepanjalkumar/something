package main

import (
	"fmt"
)

func change(str *string) {
	*str = "Change"
}

func main() {

	x := "name"
	fmt.Println(x)
	change(&x)
	fmt.Println(x)

}
