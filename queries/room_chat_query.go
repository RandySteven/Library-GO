package queries

const (
	InsertRoomChatQuery GoQuery = `
		INSERT INTO room_chats (room_name)
		VALUES (?)
	`

	SelectRoomChatQuery GoQuery = `
		SELECT id, room_name, created_at, updated_at, deleted_at
		FROM
		    room_chats
	`

	SelectRoomChatByIDQuery GoQuery = `
		SELECT id, room_name, created_at, updated_at, deleted_at
		FROM
		    room_chats
		WHERE id = ?
	`
)
