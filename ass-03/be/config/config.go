package config

import (
	"be/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "irfannmmh"
	password = "Irfan0809"
	dbPort   = "5432"
	dbname   = "db_weather"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database : ", err)
	}

	fmt.Println("connected")
	db.Debug().AutoMigrate(models.Weather{})
}

func GetDB() *gorm.DB {
	return db
}
