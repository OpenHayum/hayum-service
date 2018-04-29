package route

import "bitbucket.org/hayum/hayum-service/service"

type authRoute struct {
	router Router
	s      service.AuthServicer
}

func initAuthRoute(router Router) {
	accountService := service.NewAuthService()

	a := &authRoute{router}
}
