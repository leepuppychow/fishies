package main

import (
	"log"
	"net/http"
	"os"

	"github.com/leepuppychow/fishies/database"
	"github.com/leepuppychow/fishies/routes"
)

func main() {
	database.Connect()
	port := os.Getenv("SERVER_PORT")
	router := routes.NewRouter()

	log.Println("Server running at:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
