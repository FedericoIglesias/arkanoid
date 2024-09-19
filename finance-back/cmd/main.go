package main

import (
	"finance-api/config"
	"finance-api/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}

	os.Setenv("SECRET_KEY", "SECRET_KEY")

	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":3000")

}
