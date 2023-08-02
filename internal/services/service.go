package services

import (
	"musics_streaming_api/internal/models"
	"musics_streaming_api/internal/repository"
)

type Service interface {
	AddSong(song *models.Songs) (*models.Songs, bool)
	GetAllSongs() []*models.Songs
	UpdateSong(id string, updatedSong *models.UpdateSongs) (*models.Songs, bool)
	DeleteSong(id string) bool
}

type ServiceImpl struct {
	Repo repository.Repository
}

func ProvideService(repo repository.Repository) *ServiceImpl {
	return &ServiceImpl{
		Repo: repo,
	}
}

func (s *ServiceImpl) AddSong(song *models.Songs) (*models.Songs, bool) {
	return s.Repo.AddSong(song)
}

func (s *ServiceImpl) GetAllSongs() []*models.Songs {
	return s.Repo.GetAllSongs()
}

func (s *ServiceImpl) UpdateSong(id string, updatedSong *models.UpdateSongs) (*models.Songs, bool) {
	return s.Repo.UpdateSong(id, updatedSong)
}

func (s *ServiceImpl) DeleteSong(id string) bool {
	return s.Repo.DeleteSong(id)
}
