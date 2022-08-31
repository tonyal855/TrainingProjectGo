package main

import (
	service "TestProject/TrainingProjectGo/services"
	"net/http"
)

func main() {
	var db []*service.User
	usr := service.NewUserService(db)
	http.HandleFunc("/register", usr.RegisterHandler)
	http.HandleFunc("/getuser", usr.GetuserHandler)
	address := "localhost:8000"
	http.ListenAndServe(address, nil)

}
