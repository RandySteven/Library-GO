package responses

import "time"

type (
	CreateRoomChatResponse struct {
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	ListRoomChatsResponse struct {
		RoomChatID   uint64 `json:"room_chat_id"`
		RoomChatName string `json:"room_chat_name"`
		NumberOfUser uint64 `json:"number_of_user"`
	}

	ChatResponse struct {
		SenderID    uint64    `json:"sender_id"`
		SenderName  string    `json:"sender_name"`
		ChatMessage string    `json:"chat_message"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	RoomChatsResponse struct {
		RoomChatID    uint64          `json:"room_chat_id"`
		RoomName      string          `json:"room_name"`
		ChatResponses []*ChatResponse `json:"chat_responses"`
	}
)
