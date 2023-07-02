package chat

import (
	"fmt"
	"goilerplate/internal/domain"
	"goilerplate/pkg/server"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type handler struct {
	chatUseCase domain.ChatUseCase
}

func NewHandler(chatUseCase domain.ChatUseCase) *handler {
	return &handler{
		chatUseCase: chatUseCase,
	}
}

func (h *handler) Serve(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()

	content, err := ioutil.ReadFile("page/index.html")
	if err != nil {
		log.Println("[chat] [handler] [Serve] error: ", err)
		server.RenderError(w, http.StatusInternalServerError, err, timeStart)
		return
	}

	fmt.Fprintf(w, "%s", content)
}

func (h *handler) Upgrade(w http.ResponseWriter, r *http.Request) {
	currentGorillaConn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	username := r.URL.Query().Get("username")
	currentConn := domain.WebSocketConnection{Conn: currentGorillaConn, Username: username}
	domain.Connections = append(domain.Connections, &currentConn)

	go h.chatUseCase.HandleIO(&currentConn, domain.Connections)
}
