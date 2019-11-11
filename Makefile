# go #
GOBUILD=go build
GOCLEAN=go clean
GOTEST=go test
GOGET=go get
GORUN=go run

# git #
GIT = git
GIT_PULL_MASTER = $(GIT) pull origin master
GIT_DIFF = $(GIT) diff
GIT_ADD = $(GIT) add .
GIT_COMMIT = $(GIT) commit
GIT_PUSH = $(GIT) push

# path | binary #
CMD_PATH=./cmd/main.go
BINARY_NAME=hack-bowl

# utility #
CONFIRM = read -p  '問題があった場合は ctr + cを押してください'



help: ## コマンド一覧
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## バイナリファイルの生成(mac用)
	$(GOBUILD) $(CMD_PATH)

deploy: ## バイナリファイルの生成とデプロイ(linux用)
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) $(CMD_PATH)

clean: ## バイナリファイルの削除(go clean)
	$(GOCLEAN)
	rm -rf $(BINARY_NAME)

run: ## バイナリファイルの生成と実行
	$(GOBUILD) $(CMD_PATH)
	./main

dev: ## 開発環境としてアプリケーションを走らせる
	$(GORUN) $(CMD_PATH)

git: ## git pushまで一気にしていくれる
	$(GIT_DIFF)
	$(CONFIRM)
	$(GIT_ADD)
	$(GIT_COMMIT)
	$(GIT_PUSH)

benchmarkAuth: ## ベンチマークテスト（Auth）
	$(GOTEST) ./pkg/server -bench Auth -benchtime 5s -benchmem -cpuprofile pprof/auth.out

benchmarkUser: ## ベンチマークテスト（User）
	$(GOTEST) ./pkg/server -bench User -benchtime 5s -benchmem -cpuprofile pprof/user.out


