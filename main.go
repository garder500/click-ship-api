package main

import (
	utils "clickship/models"
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
	database, initErr = utils.InitDatabase()
	if initErr != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", initErr))
	}
}

func main() {
	// the Database successfully connected => do something now :)
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	log.Panic(http.ListenAndServe(":4000", r))
}
