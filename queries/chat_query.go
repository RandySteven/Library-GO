package queries

const (
	InsertIntoChatQuery GoQuery = `
		INSERT INTO chats (room_chat_id, user_id, chat)
		VALUES 
		    (?, ?, ?)
	`

	SelectFromChatBasedOnRoom GoQuery = `
		SELECT id, room_chat_id, user_id, chat
		FROM
		    chats
		WHERE room_chat_id = ?
	`
)
