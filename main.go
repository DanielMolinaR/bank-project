package main

import (
	"log"

	"github.com/DanielMolinaR/bank-project/server"
	"github.com/DanielMolinaR/bank-project/storage/postgres_storage"
)

func main() {
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
