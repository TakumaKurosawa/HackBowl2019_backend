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

type UserCtrl struct {
	UserSrv *service.UserService
}

func NewUserCtl(userRepo repository.UserRepository) *UserCtrl {
	return &UserCtrl{
		UserSrv: service.NewUserService(userRepo),
	}
}

// HandleUserGet ユーザ情報取得処理
func (ctrl *UserCtrl) HandleUserGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストヘッダからx-token(認証トークン)を取得
		token := request.Header.Get("x-token")
		if len(token) == 0 {
			log.Println("x-token is empty")
			return
		}

		selectedUser, err := ctrl.UserSrv.GetUserByAuthToken(token)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// レスポンスに必要な情報を詰めて返却
		response.Success(writer, responseBody.UserName{Name: selectedUser.Name})
	}
}

// HandleUserUpdate ユーザ情報更新処理
func (ctrl UserCtrl) HandleUserUpdate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// RequestBodyのパース
		var reqBody requestBody.UserUpdate
		err := json.NewDecoder(request.Body).Decode(&reqBody)
		if err != nil {
			log.Println(err)
			response.BadRequest(writer, "Invalid Request Body")
			return
		}

		// リクエストヘッダからx-token(認証トークン)を取得
		token := request.Header.Get("x-token")
		if len(token) == 0 {
			log.Println("x-token is empty")
			return
		}

		// リクエストボディのNameもEmailも空の場合はエラーを返す
		if reqBody.Name == "" && reqBody.Email == "" {
			response.InternalServerError(writer, "ユーザ情報の更新には名前かメールアドレスどちらかを含んでいる必要があります。")
			return
		}

		// どちらも変更する時
		if reqBody.Email != "" && reqBody.Name != "" {
			err = ctrl.UserSrv.ChangeUserData(token, reqBody.Name, reqBody.Email)
			if err != nil {
				log.Println(err)
				response.InternalServerError(writer, "Internal Server Error")
				return
			}
		}

		// Emailが空の時（名前だけの更新の時）
		if reqBody.Email == "" {
			err = ctrl.UserSrv.ChangeUserName(token, reqBody.Name)
			if err != nil {
				log.Println(err)
				response.InternalServerError(writer, "Internal Server Error")
				return
			}
		}

		// Nameが空の時（メールアドレスだけの更新の時）
		if reqBody.Name == "" {
			err = ctrl.UserSrv.ChangeUserEmail(token, reqBody.Email)
			if err != nil {
				log.Println(err)
				response.InternalServerError(writer, "Internal Server Error")
				return
			}
		}

		response.Success(writer, "")
	}
}
