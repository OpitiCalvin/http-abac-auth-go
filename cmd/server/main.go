package main

import (
	"fmt"

	// "log"
	// "novelsTradeIn/pkg/api"
	// "novelsTradeIn/pkg/repository"
	"os"

	"github.com/OpitiCalvin/http-abac-auth-go/pkg/server"
	"github.com/gorilla/mux"
	// "github.com/joho/godotenv"
	// "gorm.io/gorm"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

// run function is responsible for setting up db connections, routers etc
func run() error {
	// // load env params
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err) // error loading env file
	// }

	// username := os.Getenv("USERNAME")
	// password := os.Getenv("PASSWORD")
	// dbname := os.Getenv("DBNAME")
	// host := os.Getenv("HOST")
	// DBPort := os.Getenv("DBPORT")

	// connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", username, password, host, DBPort, dbname)

	// // setup a connection to database
	// db, err := setupDatabase(connectionString)
	// if err != nil {
	// 	return err
	// }

	// // create storage dependency
	// storage := repository.NewStorage(db)

	// // run migrations
	// err = storage.RunMigrations(connectionString)
	// if err != nil {
	// 	return err
	// }

	// create router dependency
	router := mux.NewRouter()
	// TO DO: implement CORS with router

	// // create user service
	// userService := api.NewUserService(storage)

	// start server
	server := server.NewServer(router)
	err := server.Run()
	if err != nil {
		return err
	}

	return nil
}

// func setupDatabase(connString string) (*gorm.DB, error) {
// 	db, err := gorm.Open("postgres", connString)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }
