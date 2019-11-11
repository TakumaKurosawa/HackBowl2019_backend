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

type AuthCtrl struct {
	AuthSrv *service.AuthService
}

func NewAuthCtl(repo repository.UserRepository) *AuthCtrl {
	return &AuthCtrl{
		AuthSrv: service.NewAuthService(repo),
	}
}

func (ctrl *AuthCtrl) HandleAuthCreate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// RequestBodyのパース
		var reqBody requestBody.Auth
		err := json.NewDecoder(request.Body).Decode(&reqBody)
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, "Invalid Request Body")
			return
		}

		authToken, err := ctrl.AuthSrv.Create(reqBody.Name, reqBody.Email, reqBody.Password)
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, "Invalid Server Error")
			return
		}

		// 生成した認証トークンを返却
		response.Success(writer, responseBody.AuthBody{Token: authToken})

	}
}
