package main

import (
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	No     int
	Name   string
	Adress string
	Job    string
	Desc   string
}

func main() {
	var getArg string = os.Args[1]
	number, _ := strconv.Atoi(getArg)

	fmt.Println(number)

	var getData = func(employees []Employee) {
		for _, people := range employees {
			if people.No == number {
				fmt.Println(people.Name, people.Adress, people.Job, people.Desc)
				break
			}
		}
	}

	var peoples = []Employee{
		{No: 1, Name: "Anto", Adress: "Jl.MNC", Job: "Dev", Desc: "test"},
		{No: 2, Name: "Jeke", Adress: "Jl.MNC", Job: "Dev", Desc: "test"},
		{No: 3, Name: "Hans", Adress: "Jl.MNC", Job: "Dev", Desc: "test"},
	}

	getData(peoples)
}
