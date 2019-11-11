package middleware

import (
	"backend_api/domain/repository"
	"backend_api/domain/service"
	"backend_api/pkg/dcontext"
	"backend_api/pkg/server/response"
	"context"
	"log"
	"net/http"
)

type AuthenticateCtrl struct {
	UserSrv *service.UserService
}

func NewAuthenticateCtl(userRepo repository.UserRepository) *AuthenticateCtrl {
	return &AuthenticateCtrl{
		UserSrv: service.NewUserService(userRepo),
	}
}

// Authenticate ユーザ認証を行ってContextへユーザID情報を保存する
func (ctrl *AuthenticateCtrl) Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		ctx := request.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		// リクエストヘッダからx-token(認証トークン)を取得
		token := request.Header.Get("x-token")
		if len(token) == 0 {
			log.Println("x-token is empty")
			return
		}

		// データベースから認証トークンに紐づくユーザの情報を取得
		user, err := ctrl.UserSrv.GetUserByAuthToken(token)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Invalid token")
			return
		}
		if user == nil {
			log.Printf("user not found. token=%s", token)
			response.BadRequest(writer, "Invalid token")
			return
		}

		// userIdをContextへ保存して以降の処理に利用する
		ctx = dcontext.SetUserID(ctx, user.AuthToken)

		// 次の処理
		nextFunc(writer, request.WithContext(ctx))
	}
}
