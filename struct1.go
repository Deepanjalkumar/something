package main

import (
	"fmt"
)

type Employee struct {
	Firstname, Lastname, Email string
	Age                        int
	MonthlySalary              []Salary
}

type Salary struct {
	Basic, HRA, TA int
}

func main() {

	e := Employee{
		Firstname: "Rana",
		Lastname:  "kumar",
		Email:     "test@test.com",
		Age:       25,
		MonthlySalary: []Salary{Salary{
			Basic: 60000,
			HRA:   10000,
			TA:    5000,
		}},
	}

	for index, _ := range e.MonthlySalary {
		fmt.Println(e.MonthlySalary[index].Basic)
	}

}
