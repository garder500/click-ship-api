package utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	dsn := "host=127.0.0.1 user=clickship password=clickship dbname=clickship port=5432 sslmode=disable TimeZone=Europe/Paris"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
