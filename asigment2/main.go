package main

import (
	"assigment2/db"
	"assigment2/server"

	"assigment2/server/controllers"
	"assigment2/server/repositories/gorm"
)

// @description    Sample API Spec for Orders
// @termsOfService http://swagger.io/terms/
// @BasePath       /
// @contact.name   Toni
// @contact.email  toni@gmail.com
func main() {
	db := db.ConnectGorm()
	ordersRepo := gorm.NewOrderRepo(db)
	ordersController := controllers.NewOrderController(ordersRepo)
	router := server.NewRouter(ordersController)
	router.Start(":9000")
}
