package requests

import "io"

type (
	CreateRoomRequest struct {
		Name      string    `form:"name"`
		Thumbnail io.Reader `form:"thumbnail"`
	}

	UploadRoomPhoto struct {
		RoomID uint64      `form:"room_id"`
		Photos []io.Reader `form:"photos"`
	}
)
