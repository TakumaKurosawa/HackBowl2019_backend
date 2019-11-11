package server

import (
	"backend_api/application/controller/requestBody"
	"backend_api/application/controller/responseBody"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

type mockResponseWriter struct {
	header http.Header
	*bytes.Buffer
}

func NewMockResponseWriter() *mockResponseWriter {
	return &mockResponseWriter{
		header: http.Header{},
		Buffer: &bytes.Buffer{},
	}
}

func (w *mockResponseWriter) Header() http.Header {
	return w.header
}

func (w *mockResponseWriter) WriteHeader(statusCode int) {

}

func Benchmark_HANDLERS_Auth(b *testing.B) {
	authHandler := InitializeAuthHandler()
	characterHandler := InitializeCharacterHandler()
	gachaHandler := InitializeGachaHandler()
	rankingHandler := InitializeRankingHandler()
	userHandler := InitializeUserHandler()
	authenticator := InitializeAuthenticator()

	reqName := requestBody.Auth{Name: "黒澤拓磨"}
	payloadName, err := json.Marshal(reqName)
	if err != nil {
		panic(err)
	}

	reqGacha := requestBody.Gacha{Times: 10}
	payloadGacha, err := json.Marshal(reqGacha)
	if err != nil {
		panic(err)
	}

	writers := NewMockResponseWriter()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		requestsAuth, _ := http.NewRequest("POST", "auth/create/", bytes.NewReader(payloadName))
		requestsChara, _ := http.NewRequest("GET", "/character/list", nil)
		requestsGacha, _ := http.NewRequest("POST", "/gacha/draw", bytes.NewReader(payloadGacha))
		requestsRankGet, _ := http.NewRequest("GET", "/ranking/get", nil)
		requestsRankList, _ := http.NewRequest("GET", "/ranking/list?start=0&end=99", nil)
		requestsUserGet, _ := http.NewRequest("GET", "/user/get", nil)
		requestsUserPost, _ := http.NewRequest("POST", "/user/update", bytes.NewReader(payloadName))

		requestsAuth.Header.Set("x-token", "dc89df59-4686-4fa9-a1ed-f5eb3a2e9c23")
		requestsAuth.Header.Set("Content-Type", "application/json")

		authenticator.Authenticate(authHandler.HandleAuthCreate())(writers, requestsAuth)
		var authResponse responseBody.AuthBody
		err := json.Unmarshal(writers.Buffer.Bytes(), &authResponse)
		if err != nil {
			log.Println(err)
		}
		writers.Buffer.Reset()

		requestsChara.Header.Set("x-token", authResponse.Token)
		requestsChara.Header.Set("Content-Type", "application/json")
		requestsGacha.Header.Set("x-token", authResponse.Token)
		requestsGacha.Header.Set("Content-Type", "application/json")
		requestsRankGet.Header.Set("x-token", authResponse.Token)
		requestsRankGet.Header.Set("Content-Type", "application/json")
		requestsRankList.Header.Set("x-token", authResponse.Token)
		requestsRankList.Header.Set("Content-Type", "application/json")
		requestsUserGet.Header.Set("x-token", authResponse.Token)
		requestsUserGet.Header.Set("Content-Type", "application/json")
		requestsUserPost.Header.Set("x-token", authResponse.Token)
		requestsUserPost.Header.Set("Content-Type", "application/json")

		authenticator.Authenticate(userHandler.HandleUserGet())(writers, requestsUserGet)
		writers.Buffer.Reset()

		authenticator.Authenticate(userHandler.HandleUserUpdate())(writers, requestsUserPost)
		writers.Buffer.Reset()

		authenticator.Authenticate(gachaHandler.HandleGachaDraw())(writers, requestsGacha)
		writers.Buffer.Reset()

		authenticator.Authenticate(characterHandler.HandleCharacterListGet())(writers, requestsChara)
		writers.Buffer.Reset()

		authenticator.Authenticate(rankingHandler.HandleUserRankingGet())(writers, requestsRankGet)
		writers.Buffer.Reset()

		authenticator.Authenticate(rankingHandler.HandleRankingListGet())(writers, requestsRankList)
		writers.Buffer.Reset()

	}
}
