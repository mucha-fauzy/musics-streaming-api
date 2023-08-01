package models

import "time"

type Songs struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Artist      string    `json:"artist"`
	Album       string    `json:"album"`
	Duration    string    `json:"duration"`
	Genre       string    `json:"genre"`
	ReleaseDate string    `json:"release_date"`
	StreamURL   string    `json:"stream_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateSongs struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Artist      string    `json:"artist"`
	Album       string    `json:"album"`
	Duration    string    `json:"duration"`
	Genre       string    `json:"genre"`
	ReleaseDate string    `json:"release_date"`
	StreamURL   string    `json:"stream_url"`
	UpdatedAt   time.Time `json:"updated_at"`
}
