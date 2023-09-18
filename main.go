package main

import (
	"basic-trade-api/database"
	"basic-trade-api/routers"
	"os"
)

func main() {
	// Start database
	database.StartDB()
	// if err != nil {
	// 	log.Fatalln("Error starting DB:", err)
	// 	return
	// }
	// Define port based on .env
	var PORT = ":" + os.Getenv("WEB_SERVER_PORT")
	// Start server
	routers.StartServer().Run(PORT)

}
