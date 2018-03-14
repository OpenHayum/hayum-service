package service

import "errors"

type S3Servicer interface {
	Upload(filename string) error
}

type s3Service struct {
}

func (s3 *s3Service) Upload(filename string) error {
	return errors.New("")
}
