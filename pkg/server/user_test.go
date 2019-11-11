package server

import (
	"backend_api/application/controller/requestBody"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func Benchmark_HANDLERS_User(b *testing.B) {
	userHandler := InitializeUserHandler()
	authenticator := InitializeAuthenticator()

	reqName := requestBody.Auth{Name: "黒澤拓磨"}
	payloadName, err := json.Marshal(reqName)
	if err != nil {
		panic(err)
	}

	writers := NewMockResponseWriter()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestsUserGet, _ := http.NewRequest("GET", "/user/get", nil)
		requestsUserPost, _ := http.NewRequest("POST", "/user/update", bytes.NewReader(payloadName))

		requestsUserGet.Header.Set("x-token", "dc89df59-4686-4fa9-a1ed-f5eb3a2e9c23")
		requestsUserGet.Header.Set("Content-Type", "application/json")
		requestsUserPost.Header.Set("x-token", "dc89df59-4686-4fa9-a1ed-f5eb3a2e9c23")
		requestsUserPost.Header.Set("Content-Type", "application/json")

		authenticator.Authenticate(userHandler.HandleUserGet())(writers, requestsUserGet)
		writers.Buffer.Reset()

		authenticator.Authenticate(userHandler.HandleUserUpdate())(writers, requestsUserPost)
		writers.Buffer.Reset()

	}
}
