package cmd

import (
	"goilerplate/config"
	"goilerplate/internal/routes"

	"github.com/gorilla/mux"
)

func initializeServer(api *mux.Router, cfg config.Config) {
	routes.InitializeApi(api, cfg)
}
