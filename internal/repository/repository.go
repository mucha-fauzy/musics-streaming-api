package repository

import "musics_streaming_api/internal/models"

type Repository interface {
	SongRepository
}

type RepositoryImpl struct {
	songs []*models.Songs
}

func ProvideRepo() *RepositoryImpl {
	return &RepositoryImpl{}
}
