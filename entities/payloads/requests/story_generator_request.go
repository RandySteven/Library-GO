package requests

type (
	StoryGeneratorRequest struct {
		Story     string `json:"story"`
		Setting   string `json:"setting"`
		Genre     string `json:"genre"`
		Style     string `json:"style"`
		Character string `json:"character"`
		Moral     string `json:"moral"`
	}
)
