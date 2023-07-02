package domain

import (
	"github.com/gorilla/websocket"
)

type ChatUseCase interface {
	HandleIO(currentConn *WebSocketConnection, connections []*WebSocketConnection)
}

type ChatRepository interface {
	EjectConnection(currentConn *WebSocketConnection)
	BroadcastMessage(currentConn *WebSocketConnection, kind, message string)
}

type SocketPayload struct {
	Message string
}

type SocketResponse struct {
	From    string
	Type    string
	Message string
}

type WebSocketConnection struct {
	*websocket.Conn
	Username string
}

const MESSAGE_NEW_USER = "New User"
const MESSAGE_CHAT = "Chat"
const MESSAGE_LEAVE = "Leave"

var Connections = make([]*WebSocketConnection, 0)
