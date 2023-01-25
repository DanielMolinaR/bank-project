package main

import (
	"log"

	"github.com/DanielMolinaR/bank-project/server"
	"github.com/DanielMolinaR/bank-project/storage/postgres_storage"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	store, err := postgres_storage.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := server.NewApiServer(":3000", store)
	server.Run()
}
