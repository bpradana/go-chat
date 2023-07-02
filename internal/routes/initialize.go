package routes

import (
	"goilerplate/config"

	"github.com/gorilla/mux"
)

func InitializeApi(api *mux.Router, cfg config.Config) {
	InitHealthRoute(api)
	InitChatRoute(api, cfg)
}
