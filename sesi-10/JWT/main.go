package main

import (
	"JWT/database"
	"JWT/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
