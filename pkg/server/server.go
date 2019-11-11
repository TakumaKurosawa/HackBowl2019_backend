package server

import (
	"log"
	"net/http"
)

// Serve HTTPサーバを起動する
func Serve(addr string) {

	authHandler := InitializeAuthHandler()
	loginHandler := InitializeLoginHandler()
	userHandler := InitializeUserHandler()
	projectHandler := InitializeProjectHandler()
	authenticator := InitializeAuthenticator()

	/* ===== URLマッピングを行う ===== */
	http.HandleFunc("/auth/create", post(authHandler.HandleAuthCreate()))

	http.HandleFunc("/login", post(loginHandler.HandleAuthenticate()))

	http.Handle("/user/get",
		get(authenticator.Authenticate(userHandler.HandleUserGet())))

	http.HandleFunc("/user/update",
		post(authenticator.Authenticate(userHandler.HandleUserUpdate())))

	http.HandleFunc("/projects",
		get(projectHandler.HandleAllProjectGet()))

	http.HandleFunc("/project", post(projectHandler.HandleProjectGet()))

	http.HandleFunc("/mail", post(authenticator.SendMail()))

	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

// get GETリクエストを処理する
func get(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodGet)
}

// post POSTリクエストを処理する
func post(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodPost)
}

// httpMethod 指定したHTTPメソッドでAPIの処理を実行する
func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// プリフライトリクエストは処理を通さない
		if request.Method == http.MethodOptions {
			return
		}
		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method Not Allowed"))
			return
		}

		// 共通のレスポンスヘッダを設定
		writer.Header().Add("Content-Type", "application/json")
		apiFunc(writer, request)
	}
}
