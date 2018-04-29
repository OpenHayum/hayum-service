package service

import (
	"mime/multipart"
	"time"

	"gopkg.in/mgo.v2/bson"

	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/core/s3"
	"bitbucket.org/hayum/hayum-service/models"

	"bitbucket.org/hayum/hayum-service/repository"
)

// S3Servicer holds the S3Service contracts
type S3Servicer interface {
	Upload(file multipart.File, header *multipart.FileHeader) (string, error)
}

// S3Service holds the S3DocumentRepository
type S3Service struct {
	repository repository.S3DocumentRepositorer
}

// NewS3DocumentService creates a new S3DocumentService
func NewS3DocumentService() *S3Service {
	return &S3Service{repository.NewS3DocumentRepository()}
}

// Upload uploads a file
func (s *S3Service) Upload(file multipart.File, header *multipart.FileHeader) (string, error) {
	s3Document := new(models.S3Document)

	s3Document.ID = bson.NewObjectId()
	s3Document.OriginalFileName = header.Filename
	s3Document.CreatedDate = time.Now()

	s3Directory := config.App.GetString("s3.directory")
	s3Bucket := config.App.GetString("s3.bucket")

	uploadManager := s3.New(s3Directory, s3Bucket)
	uploadManager.Upload(file, header.Filename)

	if err := s.repository.CreateNewS3Document(s3Document); err != nil {
		return "", err
	}

	return s3Document.ID.Hex(), nil
}
