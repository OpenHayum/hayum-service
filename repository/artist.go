package repository

import "bitbucket.org/hayum/hayum-service/models"

type ArtistRepository struct{
	*Repository
}

func NewArtistRepository() *ArtistRepository {
	return &ArtistRepository{NewRepository{}}
}

func (repository *ArtistRepository) GetArtistById(id string) (*models.Artist, error) {
	var artist models.Artist
	err := 
}
