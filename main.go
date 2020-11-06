package main

import (
	"github.com/joho/godotenv"
	"github.com/mdmims/go-azure-catalog-api/handler"
	"log"
	"os"
)

func initEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initEnv()
	a := handler.App{}

	a.Initialize(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	a.Run(":8010")
}
