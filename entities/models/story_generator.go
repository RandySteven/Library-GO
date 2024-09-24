package models

import "time"

type StoryGenerator struct {
	ID        uint64
	Prompt    string
	Result    string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
