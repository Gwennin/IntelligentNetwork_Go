package server

import (
	"log"
	"net/http"
	"os"

	"github.com/Gwennin/IntelligentNetwork_Go/src/managers"
	"github.com/Gwennin/IntelligentNetwork_Go/src/managers/db"
	"github.com/joho/godotenv"
)

func StartServer() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitializeDB()

	getRoutes()
	managers.CleanSessions()

	listenAddr := os.Getenv("LISTEN_ADDRESS")
	listenPort := os.Getenv("LISTEN_PORT")

	listenMask := listenAddr + ":" + listenPort
	log.Fatal(http.ListenAndServe(listenMask, nil))

	db.CloseDB()
}
