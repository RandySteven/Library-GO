package requests

type (
	CreateChatRoom struct {
		UserID        uint64   `json:"id"`
		InviteUserIDs []uint64 `json:"invite_user_ids"`
		RoomName      string   `json:"room_name"`
	}

	SendChat struct {
		RoomID uint64 `json:"room_id"`
		Chat   string `json:"chat"`
	}
)
