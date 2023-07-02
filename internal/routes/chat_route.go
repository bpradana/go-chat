package routes

import (
	"goilerplate/config"
	chat "goilerplate/internal/chat"

	"github.com/gorilla/mux"
)

func InitChatRoute(api *mux.Router, cfg config.Config) {
	userRepository := chat.NewRepository()
	userUseCase := chat.NewUseCase(userRepository)
	userHandler := chat.NewHandler(userUseCase)

	api.HandleFunc("/", userHandler.Serve)
	api.HandleFunc("/ws", userHandler.Upgrade)
}
