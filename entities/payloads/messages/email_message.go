package messages

type EmailMessage struct {
	ID      string `json:"id"`
	To      string `json:"to"`
	Content any    `json:"content"`
}
