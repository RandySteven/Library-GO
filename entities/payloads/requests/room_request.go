package requests

import "io"

type CreateRoomRequest struct {
	Name      string    `form:"name"`
	Thumbnail io.Reader `form:"thumbnail"`
}
