package main

import (
	utils "clickship/models"
	"fmt"

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
	// Do something with the database
	fmt.Println(database.Config.Dialector.Name())
}
