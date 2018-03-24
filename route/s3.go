package route

import (
	"fmt"
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

func initS3Route(router Router) {
	service := service.NewS3DocumentService(repository.NewRepository(router.GetMongo(), config.CollectionS3Document))
	s3 := &s3Route{router, service}

	s3.router.POST("/file/upload", s3.upload)
}

func (s3 *s3Route) upload(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := s3.service.Upload(ps.ByName("filename"))
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		s3.router.JSON(w, "")
		return
	}
	fmt.Fprintf(w, "%v", header.Header)
	s3.service.Upload(file, header)

	s3.router.JSON(w, "")
}
