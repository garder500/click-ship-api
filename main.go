package main

import (
	"clickship/models"
	"fmt"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	initErr  error
)

func init() {
	database, initErr = models.InitDatabase()
	if initErr != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", initErr))
	}

	// create Schema

	database.Exec("CREATE SCHEMA IF NOT EXISTS users")
	database.Exec("CREATE SCHEMA IF NOT EXISTS suscribtions")

	// Migrate the schema
	if database.AutoMigrate(&models.User{}, &models.UserAdress{}, &models.UserCard{}, &models.UserIdentity{}, &models.Suscribtion{}, &models.SuscribtionHistory{}) != nil {
		panic("Failed to migrate the schema")
	} else {
		log.Println("Schema migrated successfully")
	}
}

func main() {
	// the Database successfully connected => do something now :)
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	log.Println("Server is running on port 4000")
	log.Panic(http.ListenAndServe(":4000", r))
}
