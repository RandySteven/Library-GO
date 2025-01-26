package responses

import "time"

type (
	CreateRoomResponse struct {
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	ListRoomResponse struct {
		ID          uint64     `json:"id"`
		RoomName    string     `json:"room_name"`
		Thumbnail   string     `json:"thumbnail"`
		IsAvailable bool       `json:"is_available"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at"`
		DeletedAt   *time.Time `json:"deleted_at"`
	}

	RoomDetailResponse struct {
		ID          uint64     `json:"id"`
		RoomName    string     `json:"room_name"`
		Thumbnail   string     `json:"thumbnail"`
		IsAvailable bool       `json:"is_available"`
		Photos      []string   `json:"photos"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at"`
		DeletedAt   *time.Time `json:"deleted_at"`
	}

	UploadRoomPhotoResponse struct {
		RoomID    uint64   `json:"room_id"`
		RoomName  string   `json:"room_name"`
		PhotoURLs []string `json:"photo_urls"`
	}
)
