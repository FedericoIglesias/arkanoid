package main

import (
	"finance-api/config"
	"finance-api/routes"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":3000")
}
