package main

import (
	"fmt"
)

type Users []User

type User struct {
	Firstname, Lastname, Email string
}

func main() {

	var sub Users

	sub = append(sub, User{
		Firstname: "Rana",
		Lastname:  "Kumar",
		Email:     "test@test.com",
	}, User{
		Firstname: "Ram",
		Lastname:  "kumar",
		Email:     "test@test.com",
	})

	for index, _ := range sub {
		fmt.Println(sub[index].Firstname)
	}

}
