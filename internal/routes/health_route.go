package routes

import (
	"goilerplate/pkg/server"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func InitHealthRoute(api *mux.Router) {
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		resp := "service healthy"

		server.RenderResponse(w, http.StatusOK, resp, timeStart)
	}).Methods(http.MethodGet)
}
