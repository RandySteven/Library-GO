package models

import "time"

type Event struct {
	ID                        uint64
	Title                     string
	Thumbnail                 string
	Price                     *uint64
	Description               string
	ParticipantNumber         uint64
	OccupiedParticipantNumber uint64
	Date                      time.Time
	StartTime                 time.Time
	EndTime                   time.Time
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	DeletedAt                 *time.Time
}
