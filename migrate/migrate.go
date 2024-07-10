package main

import (
	"uniwise/animals/initializers"
	"uniwise/animals/models"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Animal{})
}
