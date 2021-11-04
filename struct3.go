package main

import (
	"encoding/json"
	"fmt"
)

type Users1 []User1

type User1 struct {
	Name   string `json:"Name"`
	Age    string `json:"Age"`
	Social string `json:"Social"`
}

func main() {

	var sub Users

	sub = append(sub, User1{
		Name:   "Rana",
		Age:    25,
		Social: "https://www.facebook.com/rana",
	},
		User1{
			Name:   "Ram",
			Age:    25,
			Social: "https://www.facebook.com/rama"
		})

	file_byte, _ := json.Marshal(sub)
	fmt.Println(string(file_byte))

}
