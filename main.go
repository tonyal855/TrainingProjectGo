package main

import "fmt"

type Person struct {
	name string
}

func main() {
	var NameF = func(persons []*Person) []*Person {
		res := []*Person{}
		for _, data := range persons {
			fmt.Println(data.name)
			res = append(res, data)
		}
		return res
	}

	var anto = Person{name: "anto"}
	var rara = Person{name: "rara"}

	list_name := []*Person{&anto, &rara}
	fmt.Println(NameF(list_name))
}
