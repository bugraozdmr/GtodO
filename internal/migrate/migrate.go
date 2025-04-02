package main

import (
	app "gtodo/internal/app/entity"
	"gtodo/internal/config"
	db "gtodo/internal/database"
	"log"
)

func main() {
	cnf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load environements")
	}
	database := db.ConnectPGDB(cnf)

	if err := database.AutoMigrate(
		&app.UserRegister{},
		&app.Todo{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Migration successful")
}
