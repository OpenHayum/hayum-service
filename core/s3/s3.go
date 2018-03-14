package s3

import "errors"

type s3Manager struct {
}

func New() *s3Manager {
	return &s3Manager{}
}

func (s *s3Manager) Upload(content []byte) error {
	return errors.New("")
}

func (s *s3Manager) Download(filename string) []byte {
	return []byte{1, 2}
}
