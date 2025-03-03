package main

import (
	"log"
	"os"

	"github.com/flyluman/portfolio-on-golang/internal/database"
	"github.com/flyluman/portfolio-on-golang/internal/local_time"
	"github.com/flyluman/portfolio-on-golang/internal/services"
)

func main() {
	local_time.InitTime()

	database.InitDB()
	defer database.DB.Close()

	server := services.NewAPIServer(os.Getenv("PORT"), os.Getenv("CERTFILE"), os.Getenv("KEYFILE"))
	log.Fatal(server.Run())
}
