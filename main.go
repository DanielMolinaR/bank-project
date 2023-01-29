package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/DanielMolinaR/bank-project/scripts/seed"
	"github.com/DanielMolinaR/bank-project/server"
	"github.com/DanielMolinaR/bank-project/storage/postgres_storage"
	"github.com/joho/godotenv"
)

func main() {
	seedDB := flag.Bool("seed", false, "See the DB")
	flag.Parse()

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

	if *seedDB {
		fmt.Println("Seeding the databse")
		seed.SeedDB(store)
	}

	server := server.NewApiServer(":3000", store)
	server.Run()
}
