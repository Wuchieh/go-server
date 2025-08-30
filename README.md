# Go Server 專案

這是一個使用 Go 1.24.6 的後端服務專案，支援 PostgreSQL、MongoDB 和 Redis。

## 專案結構

```
.
├── cmd/                  # 主程式入口
├── internal/             # 內部包
├── main.go               # 主程式
├── go.mod                # Go 模組文件
├── go.sum                # Go 依賴校驗文件
├── Makefile              # 專案構建和運行腳本
├── Dockerfile            # Docker 映像構建文件
├── docker-compose.yml    # 資料庫服務配置
├── .dockerignore         # Docker 忽略文件
└── README.md             # 專案說明文件
```

## 快速開始

### 1. 安裝依賴

```bash
make install
```

### 2. 啟動資料庫服務

```bash
make db-up
```

這會啟動以下服務：
- PostgreSQL (端口: 5432)
- MongoDB (端口: 27017)
- Redis (端口: 6379)
- pgAdmin (端口: 5050) - PostgreSQL 管理工具
- Mongo Express (端口: 8081) - MongoDB 管理工具

### 3. 運行應用

```bash
make run
```

或者使用開發模式（自動啟動資料庫和應用）：

```bash
make dev
```

## 可用的 Make 命令

### 基本命令
- `make help` - 顯示所有可用命令
- `make build` - 編譯專案
- `make run` - 運行專案
- `make test` - 運行測試
- `make clean` - 清理編譯產物

### 依賴管理
- `make deps` - 下載依賴
- `make deps-tidy` - 整理依賴

### Docker 相關
- `make docker-build` - 構建 Docker 映像
- `make docker-run` - 運行 Docker 容器
- `make docker-stop` - 停止 Docker 容器
- `make docker-clean` - 清理 Docker 映像

### 資料庫相關
- `make db-up` - 啟動資料庫服務
- `make db-down` - 停止資料庫服務
- `make db-logs` - 查看資料庫日誌

### 開發工具
- `make lint` - 代碼檢查
- `make fmt` - 格式化代碼
- `make docs` - 生成 API 文檔

## 資料庫連接資訊

### PostgreSQL
- 主機: localhost:5432
- 資料庫: go_server
- 用戶名: postgres
- 密碼: postgres123

### MongoDB
- 主機: localhost:27017
- 資料庫: go_server
- 用戶名: admin
- 密碼: admin123

### Redis
- 主機: localhost:6379
- 密碼: redis123

## 管理工具

### pgAdmin (PostgreSQL)
- 網址: http://localhost:5050
- 用戶名: admin@example.com
- 密碼: admin123

### Mongo Express (MongoDB)
- 網址: http://localhost:8081
- 用戶名: admin
- 密碼: admin123

## 環境變數配置

創建 `.env` 文件並配置以下環境變數：

```bash
# 應用配置
APP_NAME=go-server
APP_PORT=8080
APP_ENV=development

# PostgreSQL 配置
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DB=go_server
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres123

# MongoDB 配置
MONGO_URI=mongodb://admin:admin123@localhost:27017/go_server?authSource=admin

# Redis 配置
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=redis123
```

## Docker 部署

### 構建映像
```bash
make docker-build
```

### 運行容器
```bash
make docker-run
```

### 停止容器
```bash
make docker-stop
```

## 注意事項

1. 確保 Docker 和 Docker Compose 已安裝
2. 首次運行時，資料庫初始化可能需要一些時間
3. 生產環境中請修改預設密碼
4. 建議將 `.env` 文件添加到 `.gitignore` 中

## 故障排除

### 資料庫連接問題
```bash
# 檢查服務狀態
docker-compose ps

# 查看日誌
make db-logs

# 重啟服務
make db-down
make db-up
```

### 端口衝突
如果遇到端口衝突，可以在 `docker-compose.yml` 中修改端口映射。

## 貢獻

歡迎提交 Issue 和 Pull Request！
