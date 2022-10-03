package db

import (
	"assigment2/server/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host   = "localhost"
	port   = "5432"
	user   = "postgres"
	pass   = "password"
	dbname = "assigment2"
)

func ConnectGorm() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.Orders{})
	db.Debug().AutoMigrate(models.Items{})

	return db

}
