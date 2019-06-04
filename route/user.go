package route

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"hayum/core_apis/errors"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"hayum/core_apis/repository"
	"hayum/core_apis/service"
	"hayum/core_apis/util"
	"net/http"
)

type userRoute struct {
	router  Router
	service service.UserService
}

func initUserRoute(router Router) {
	userRepo := repository.NewSQLUserRepository(router.GetConn())

	userService := service.NewUserService(userRepo)
	u := &userRoute{router, userService}

	u.router.POST("/user", u.createUser)
	u.router.GET("/user/:id", u.getUser)
}

func (u *userRoute) createUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user *models.User
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&user)
	if errors.CheckAndSendResponseBadRequest(err, w) {
		return
	}
	logger.Log.Info("uuu", user)

	createdUser, err := u.service.GetByMobileOrEmail(ctx, user.Email, user.Mobile)
	if errors.CheckAndSendResponseInternalServerError(err, w) {
		return
	}
	if *createdUser != (models.User{}) {
		if errors.CheckAndSendResponseErrorWithStatus(errors.ErrUserMobileOrEmailAlreadyAssociated, w, http.StatusConflict) {
			return
		}
	}

	err = u.service.Save(ctx, user)
	if errors.CheckAndSendResponseInternalServerError(err, w) {
		return
	}

	u.router.JSONWithStatus(w, http.StatusCreated, user)
}

func (u *userRoute) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	id, err := util.StrToInt64(ps.ByName("id"))
	if errors.CheckAndSendResponseBadRequest(err, w) {
		return
	}

	user, err := u.service.GetByID(ctx, id)

	if err != nil {
		logger.Log.Info("Cannot get user by id:", id)
		http.Error(w, fmt.Sprintf("Cannot get user by id: %d", id), http.StatusNotFound)
		return
	}

	u.router.JSON(w, user)
}
