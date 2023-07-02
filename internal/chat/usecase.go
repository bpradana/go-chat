package chat

import (
	"fmt"
	"goilerplate/internal/domain"
	"log"
	"strings"
)

type usecase struct {
	chatRepository domain.ChatRepository
}

func NewUseCase(chatRepository domain.ChatRepository) domain.ChatUseCase {
	return &usecase{
		chatRepository: chatRepository,
	}
}

func (u *usecase) HandleIO(currentConn *domain.WebSocketConnection, connections []*domain.WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()

	u.chatRepository.BroadcastMessage(currentConn, domain.MESSAGE_NEW_USER, "")

	for {
		payload := domain.SocketPayload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				u.chatRepository.BroadcastMessage(currentConn, domain.MESSAGE_LEAVE, "")
				u.chatRepository.EjectConnection(currentConn)
				return
			}

			log.Println("ERROR", err.Error())
			continue
		}

		u.chatRepository.BroadcastMessage(currentConn, domain.MESSAGE_CHAT, payload.Message)
	}
}
