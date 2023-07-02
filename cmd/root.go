package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"goilerplate/config"
)

func Start() {
	// load env file to system
	err := godotenv.Load()
	if err != nil {
		log.Println("[cmd] [root] [Start] error loading .env file, err: ", err.Error())
	}

	// initialize config
	cfg := config.Initialize()

	// create mux router instance
	server := mux.NewRouter()

	// init routes
	api := server.PathPrefix("/").Subrouter()
	initializeServer(api, cfg)

	// start server
	port := os.Getenv("PORT")
	log.Printf("server started at port %s", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), server)
	if err != nil {
		log.Println("[cmd] [root] [Start] server error, err: ", err.Error())
	}
}
