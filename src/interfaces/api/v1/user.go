package apiV1

import (
	"encoding/json"
	"io/ioutil"
	"mongodb-example/src/interfaces/api/utils"
	"mongodb-example/src/usecases"
	"mongodb-example/src/usecases/dto"
	"net/http"

	"github.com/gorilla/mux"
)

type userController struct {
	*mux.Router
	logger      usecases.Logger
	userService usecases.IUserService
}

func NewUserHandler(logger usecases.Logger, userService usecases.IUserService) http.Handler {
	r := &userController{
		mux.NewRouter(),
		logger,
		userService,
	}

	r.Router.Path("/").Methods(http.MethodPost).Headers(utils.ContentType, utils.ApplicationJson).HandlerFunc(r.jsonCreateUser)

	return r
}

func (c userController) jsonCreateUser(w http.ResponseWriter, r *http.Request) {
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.logger.Errorln(err)
		if err := utils.NewInternalServerError().WriteJson(w); err != nil {
			c.logger.Errorln(err)
		}
		return
	}

	var user dto.User
	err = json.Unmarshal(raw, &user)
	if err != nil {
		c.logger.Errorln(err)
		if err := utils.NewInternalServerError().WriteJson(w); err != nil {
			c.logger.Errorln(err)
		}
		return
	}

	id, err := c.userService.Store(user)
	if err != nil {
		c.logger.Errorln(err)
		if err := utils.NewInternalServerError().WriteJson(w); err != nil {
			c.logger.Errorln(err)
		}
		return
	}

	w.Header().Set(utils.ContentType, utils.ApplicationJson)
	w.WriteHeader(http.StatusCreated)
	rawResp, err := json.Marshal(dto.Id{id})
	if _, err = w.Write(rawResp); err != nil {
		c.logger.Errorln(err)
	}
}
