package queries

const (
	InsertRoomQuery GoQuery = `
		INSERT INTO rooms (name, thumbnail, is_available)
		VALUES 
		    (?, ?, true)
	`
)
