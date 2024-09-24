package queries

const (
	InsertStoryGenerator GoQuery = `
		INSERT INTO story_generators(prompt, result, image)
		VALUES (?, ?, ?)
	`
)
