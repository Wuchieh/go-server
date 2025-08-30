# Go 專案 Makefile
.PHONY: help build run test clean docker-build docker-run docker-stop

# 變數定義
BINARY_NAME=go-server
DOCKER_IMAGE=go-server
DOCKER_TAG=latest

# 預設目標
.DEFAULT_GOAL := help

help: ## 顯示幫助信息
	@echo "可用的命令:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Go 相關命令
build: ## 編譯 Go 專案
	@echo "編譯 Go 專案..."
	@go build -o bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME)
	@echo "編譯完成: bin/$(BINARY_NAME)"

run: ## 運行 Go 專案
	@echo "運行 Go 專案..."
	@go run ./cmd/$(BINARY_NAME)

test: ## 運行測試
	@echo "運行測試..."
	@go test -v ./...

test-coverage: ## 運行測試並生成覆蓋率報告
	@echo "運行測試並生成覆蓋率報告..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "覆蓋率報告已生成: coverage.html"

clean: ## 清理編譯產物
	@echo "清理編譯產物..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@echo "清理完成"

# 依賴管理
deps: ## 下載依賴
	@echo "下載 Go 依賴..."
	@go mod download

deps-tidy: ## 整理依賴
	@echo "整理 Go 依賴..."
	@go mod tidy

# Docker 相關命令
docker-build: ## 構建 Docker 映像
	@echo "構建 Docker 映像..."
	@docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
	@echo "Docker 映像構建完成: $(DOCKER_IMAGE):$(DOCKER_TAG)"

docker-run: ## 運行 Docker 容器
	@echo "運行 Docker 容器..."
	@docker run -d --name $(BINARY_NAME) -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)
	@echo "容器已啟動，訪問 http://localhost:8080"

docker-stop: ## 停止 Docker 容器
	@echo "停止 Docker 容器..."
	@docker stop $(BINARY_NAME) || true
	@docker rm $(BINARY_NAME) || true
	@echo "容器已停止並移除"

docker-clean: ## 清理 Docker 映像
	@echo "清理 Docker 映像..."
	@docker rmi $(DOCKER_IMAGE):$(DOCKER_TAG) || true
	@echo "Docker 映像已清理"

# 資料庫相關命令
db-up: ## 啟動資料庫服務
	@echo "啟動資料庫服務..."
	@docker-compose up -d postgres mongodb redis
	@echo "資料庫服務已啟動"

db-down: ## 停止資料庫服務
	@echo "停止資料庫服務..."
	@docker-compose down
	@echo "資料庫服務已停止"

db-logs: ## 查看資料庫日誌
	@docker-compose logs -f

# 開發環境
dev: ## 啟動開發環境（資料庫 + 應用）
	@echo "啟動開發環境..."
	@make db-up
	@echo "等待資料庫啟動..."
	@sleep 10
	@make run

# 完整清理
clean-all: clean docker-clean ## 完整清理（編譯產物 + Docker）
	@echo "完整清理完成"

# 安裝依賴
install: deps ## 安裝專案依賴
	@echo "依賴安裝完成"

# 檢查代碼
lint: ## 運行代碼檢查
	@echo "運行代碼檢查..."
	@go vet ./...
	@golangci-lint run

# 格式化代碼
fmt: ## 格式化代碼
	@echo "格式化代碼..."
	@go fmt ./...
	@goimports -w .

# 生成文檔
docs: ## 生成 API 文檔
	@echo "生成 API 文檔..."
	@swag init -g cmd/$(BINARY_NAME)/main.go
	@echo "API 文檔已生成"
