package service

import (
	"errors"
	"mime/multipart"
	"time"

	"gopkg.in/mgo.v2/bson"

	"bitbucket.org/hayum/hayum-service/core/s3"
	"bitbucket.org/hayum/hayum-service/models"

	"bitbucket.org/hayum/hayum-service/repository"
)

// S3Servicer holds the S3Service contracts
type S3Servicer interface {
	Upload(file multipart.File, header *multipart.FileHeader) error
}

// S3Service holds the S3DocumentRepository
type S3Service struct {
	repository *repository.S3DocumentRepository
}

// NewS3DocumentService creates a new S3DocumentService
func NewS3DocumentService(r *repository.MongoRepository) *S3Service {
	return &S3Service{repository.NewS3DocumentRepository(r)}
}

// Upload uploads a file
func (s *S3Service) Upload(file multipart.File, header *multipart.FileHeader) error {
	s3Document := new(models.S3Document)

	s3Document.ID = bson.NewObjectId()
	s3Document.OriginalFileName = header.Filename
	s3Document.CreatedDate = time.Now()

	uploadManager := s3.New("", "")
	uploadManager.Upload(file, "")

	s.repository.CreateNewS3Document(s3Document)

	return errors.New("")
}
