package main

import "fmt"

func main() {

	var NameF = func(name []*string) []*string {
		res := []*string{}
		for _, data := range name {
			fmt.Println(data)
			res = append(res, data)
		}
		return res
	}

	anto := "anto"
	rara := "rara"
	list_name := []*string{&anto, &rara}
	fmt.Println(NameF(list_name))
}

func print_names() {
	var names = []string{"Toni", "Jaka", "Hans", "Fajar", "Stepanus", "Edwin", "Wicaksana", "Kadek", "Zakaria", "Satrio"}
	for i, data := range names {
		fmt.Println(i, data)
	}
}
