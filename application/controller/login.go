package controller

import (
	"backend_api/application/controller/requestBody"
	"backend_api/application/controller/responseBody"
	"backend_api/domain/repository"
	"backend_api/domain/service"
	"backend_api/pkg/server/response"
	"encoding/json"
	"log"
	"net/http"
)

type LoginCtrl struct {
	LoginSrv *service.LoginService
}

func NewLoginCtl(repo repository.UserRepository) *LoginCtrl {
	return &LoginCtrl{
		LoginSrv: service.NewLoginService(repo),
	}
}

func (ctrl *LoginCtrl) HandleAuthenticate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// RequestBodyのパース
		var reqBody requestBody.Login
		err := json.NewDecoder(request.Body).Decode(&reqBody)
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, "Invalid Request Body")
			return
		}

		authedUser, err := ctrl.LoginSrv.Authenticate(reqBody.Email, reqBody.Password)
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, "Invalid Server Error")
			return
		}

		// 生成した認証トークンを返却
		response.Success(writer, responseBody.LoginBody{Token: authedUser.AuthToken})

	}
}
