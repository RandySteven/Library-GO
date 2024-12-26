package requests

type (
	CreateChatRoom struct {
		InviteUserIDs []uint64 `json:"invite_user_ids"`
		RoomName      string   `json:"room_name"`
	}

	InviteUsers struct {
		UserIDs    []uint64 `json:"user_ids"`
		RoomChatID uint64   `json:"room_chat_id"`
	}

	SendChat struct {
		RoomID uint64 `json:"room_id"`
		Chat   string `json:"chat"`
	}
)
