package queries

const (
	InsertIntoRoomChatUsersQuery GoQuery = `
		INSERT INTO room_chat_users (room_chat_id, user_id)
		VALUES 
		    (?, ?)
	`

	SelectRoomChatUsersQuery GoQuery = `
		SELECT id, room_chat_id, user_id, created_at, updated_at, deleted_at
		FROM
		    room_chat_users
	`

	SelectRoomChatUsersByIDQuery GoQuery = `
		SELECT id, room_chat_id, user_id, created_at, updated_at, deleted_at
			FROM
				room_chat_users
		WHERE id = ?
	`

	SelectRoomChatUsersByRoomChatIDQuery GoQuery = `
		SELECT id, room_chat_id, user_id, created_at, updated_at, deleted_at
			FROM
				room_chat_users
		WHERE room_chat_id = ?
	`
	SelectRoomChatUsersByUserIDQuery GoQuery = `
		SELECT id, room_chat_id, user_id, created_at, updated_at, deleted_at
			FROM
				room_chat_users
		WHERE user_id = ?
	`

	CountUsersBasedOnRoomChatIDQuery GoQuery = `
		SELECT COUNT(user_id)
		FROM
		    room_chat_users
		WHERE room_chat_id = ?
		GROUP BY room_chat_id
	`
)
