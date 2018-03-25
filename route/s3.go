package route

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/repository"

	"bitbucket.org/hayum/hayum-service/service"
	"github.com/julienschmidt/httprouter"
)

type s3Route struct {
	router  Router
	service service.S3Servicer
}

type uploadResponse struct {
	DocumentID string `json:"documentID"`
}

func initS3Route(router Router) {
	service := service.NewS3DocumentService(repository.NewRepository(router.GetMongo(), config.CollectionS3Document))
	s3 := &s3Route{router, service}

	s3.router.POST("/file/upload", s3.upload)
}

func (s3 *s3Route) upload(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	file, header, err := r.FormFile("file")

	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to read multipart file", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%v", header.Header)

	docID, err := s3.service.Upload(file, header)

	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to upload file", http.StatusBadRequest)
		return
	}

	s3.router.JSON(w, &uploadResponse{DocumentID: docID})
}
