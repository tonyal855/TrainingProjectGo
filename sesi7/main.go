package main

import (
	configs "sesi7/configs"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	username = "postgres"
// 	password = "password"
// 	port     = "5432"
// 	dbname   = "training-golang"
// )

// func ConnectPosgree() (*sql.DB, error) {
// 	dsn := fmt.Sprintf(
// 		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, username, password, dbname,
// 	)

// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("SUCCESS CONNECT")
// 	return db, nil
// }

func main() {
	db, err := configs.ConnectPosgree()
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db not connect")
	}
}
