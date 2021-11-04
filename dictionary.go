package main

import (
	"fmt"
)

func main() {

	dictionary := make(map[string]int)

	dictionary["Page"] = 1

	fmt.Println(dictionary["Page"])
}
