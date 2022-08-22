package main

import "fmt"

func main() {
	var names = []string{"Toni", "Jaka", "Hans", "Fajar", "Stepanus", "Edwin", "Wicaksana", "Kadek", "Zakaria", "Satrio"}
	for i, data := range names {
		fmt.Println(i, data)
	}
}
