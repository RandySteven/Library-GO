package requests

type (
	StoryGeneratorRequest struct {
		Title         string `json:"title"`
		Setting       string `json:"setting"`
		Genre         string `json:"genre"`
		Style         string `json:"style"`
		MainCharacter string `json:"main_character"`
		SideCharacter string `json:"side_character"`
		Ending        string `json:"ending"`
	}
)
