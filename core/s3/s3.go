package s3

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// HayumS3Manager stores aws session info and s3 details
type HayumS3Manager struct {
	Session   *session.Session
	Directory string
	Bucket    string
}

// New creates a new Hayum S3 Manager
func New(directory string, bucket string) *HayumS3Manager {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")})
	if err != nil {
		log.Panic(err)
	}
	return &HayumS3Manager{sess, directory, bucket}
}

// Upload uploads the file to s3
func (s *HayumS3Manager) Upload(file multipart.File, key string) error {
	defer file.Close()
	uploader := s3manager.NewUploader(s.Session)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
		Body:   file,
	})

	fmt.Printf("File uploaded to, %s\n", aws.StringValue(&result.Location))

	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}

	return nil
}

// Download download the file from s3
func (s *HayumS3Manager) Download(filename string, key string) error {
	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(s.Session)

	// Create a file to write the S3 Object contents to.
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %q, %v", filename, err)
	}

	// Write the contents of S3 Object to the file
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
	})

	fmt.Printf("file downloaded, %d bytes\n", n)

	if err != nil {
		return fmt.Errorf("failed to download file, %v", err)
	}
	return nil
}
