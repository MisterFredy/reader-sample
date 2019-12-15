package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/reader/config"
	. "github.com/reader/core/database"

	"github.com/joho/godotenv"
)

var database = Server{}
var network = Network{}

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("getting the env values")
	}
	database.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	network.Port = os.Getenv("APP_PORT")
}

func main() {
	network.Route()
}
