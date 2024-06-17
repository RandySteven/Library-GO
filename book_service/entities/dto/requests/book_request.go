package requests

type (
	BookRequest struct {
		Title       string `json:"title"`
		Author      string `json:"author"`
		Description string `json:"description"`
		ImageURL    string `json:"image_url"`
	}
)
