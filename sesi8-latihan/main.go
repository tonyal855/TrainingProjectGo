package main

import (
	"sesi8-latihan/db"
	"sesi8-latihan/server"
	"sesi8-latihan/server/controllers"
	"sesi8-latihan/server/repositories/gorm"
)

// @description    Sample API Spec for Orders
// @termsOfService http://swagger.io/terms/
// @BasePath       /
// @contact.name   Toni
// @contact.email  toni@gmail.com
func main() {
	db := db.ConnectGorm()
	peopleRepo := gorm.NewPersonRepo(db)
	peopleController := controllers.NewPersonController(peopleRepo)
	router := server.NewRouter(peopleController)
	router.Start(":4000")
}
