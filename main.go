package main

import (
	service "TestProject/TrainingProjectGo/services"
	"fmt"
	"sync"
)

func main() {
	var db []*service.User
	usr := service.NewUserService(db)
	res := usr.Regis(&service.User{Nama: "Budi"})
	fmt.Println(res)

	res2 := usr.Regis(&service.User{Nama: "Jamal"})
	fmt.Println(res2)

	getUser := usr.GetUser()
	fmt.Print(getUser)
	var wg sync.WaitGroup
	for i, data := range getUser {
		wg.Add(1)
		go CetakNama(i, data.Nama, &wg)
	}
	wg.Wait()
}

func CetakNama(index int, nama string, wg *sync.WaitGroup) {
	fmt.Println(nama, index)
	wg.Done()
}
