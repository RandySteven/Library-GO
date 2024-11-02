package models

import "time"

type Event struct {
	ID                        uint64
	Title                     string
	Thumbnail                 string
	Description               string
	ParticipantNumber         uint64
	OccupiedParticipantNumber uint64
	Date                      time.Time
	StartDuration             time.Time
	EndDuration               time.Time
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	DeletedAt                 *time.Time
}
