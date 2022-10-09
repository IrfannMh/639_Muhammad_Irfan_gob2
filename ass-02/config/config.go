package config

import (
	"ass-02/structs"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	dsn := "host=localhost user=irfannmmh password=Irfan0809 dbname=ass02 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed connect db")
	}

	db.AutoMigrate(&structs.Order{}, &structs.Item{})
	fmt.Println("database connected")
	return db
}
