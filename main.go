package main

import (
	service "TestProject/TrainingProjectGo/services"
	"fmt"
)

func main() {
	var db []*service.User
	usr := service.NewUserService(db)
	res := usr.Regis(&service.User{Nama: "Budi"})
	fmt.Println(res)

	res2 := usr.Regis(&service.User{Nama: "Jamal"})
	fmt.Println(res2)

	getUser := usr.GetUser()

	for i, data := range getUser {
		fmt.Println(i, data)
	}
}
