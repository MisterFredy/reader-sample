package main

import (
	"fmt"
	"log"
	"os"

	"github.com/reader/core/route"

	"github.com/joho/godotenv"
	"github.com/reader/core/database"
	"github.com/reader/database/migration"
	seeder "github.com/reader/database/seeder"
)

func init() {
	var err error
	var database = database.Server{}
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("getting the env values")
	}
	database.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	//network.Port = os.Getenv("APP_PORT")
	migration.Migrate(database.DB)
	seeder.Seedload(database.DB)
}

func main() {
	//route.Run(os.Getenv("APP_PORT"))
	port := ":" + os.Getenv("APP_PORT")

	route.Run(port)
}
