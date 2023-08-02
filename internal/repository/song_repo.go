package repository

import (
	"musics_streaming_api/internal/models"
	"time"

	"github.com/google/uuid"
)

type SongRepository interface {
	AddSong(song *models.Songs) (*models.Songs, bool)
	GetAllSongs() []*models.Songs
	UpdateSong(id string, updatedSong *models.UpdateSongs) (*models.Songs, bool)
	DeleteSong(id string) bool
}

func (r *RepositoryImpl) AddSong(song *models.Songs) (*models.Songs, bool) {
	// Generate a new random UUID for the song
	song.ID = uuid.New().String()
	song.CreatedAt = time.Now()
	song.UpdatedAt = time.Now()
	r.songs = append(r.songs, song)
	return song, true
}

func (r *RepositoryImpl) GetAllSongs() []*models.Songs {
	return r.songs
}

func (r *RepositoryImpl) UpdateSong(id string, updatedSong *models.UpdateSongs) (*models.Songs, bool) {
	for _, song := range r.songs {
		if song.ID == id {
			song.Title = updatedSong.Title
			song.Artist = updatedSong.Artist
			song.Album = updatedSong.Album
			song.Duration = updatedSong.Duration
			song.Genre = updatedSong.Genre
			song.ReleaseDate = updatedSong.ReleaseDate
			song.StreamURL = updatedSong.StreamURL
			song.UpdatedAt = time.Now()
			return song, true
		}
	}
	return nil, false
}

func (r *RepositoryImpl) DeleteSong(id string) bool {
	for i, song := range r.songs {
		if song.ID == id {
			// Remove the song from the slice
			r.songs = append(r.songs[:i], r.songs[i+1:]...)
			return true
		}
	}
	return false
}
