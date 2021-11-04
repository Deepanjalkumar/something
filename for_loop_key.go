package main

import (
	"fmt"
)

func main() {

	dict := make(map[string]int)
	dict["Page"] = 1

	for key, _ := range dict {
		fmt.Println(dict[key])
	}
}
