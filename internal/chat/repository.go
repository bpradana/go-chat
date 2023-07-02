package chat

import (
	"goilerplate/internal/domain"

	"github.com/novalagung/gubrak/v2"
)

type repository struct{}

func NewRepository() domain.ChatRepository {
	return &repository{}
}

func (r *repository) EjectConnection(currentConn *domain.WebSocketConnection) {
	filtered := gubrak.From(domain.Connections).Reject(func(each *domain.WebSocketConnection) bool {
		return each == currentConn
	}).Result()
	domain.Connections = filtered.([]*domain.WebSocketConnection)
}

func (r *repository) BroadcastMessage(currentConn *domain.WebSocketConnection, kind, message string) {
	for _, eachConn := range domain.Connections {
		if eachConn == currentConn {
			continue
		}

		eachConn.WriteJSON(domain.SocketResponse{
			From:    currentConn.Username,
			Type:    kind,
			Message: message,
		})
	}
}
