# 多階段構建 Dockerfile
FROM golang:1.24.6-alpine AS builder

# 安裝必要的系統依賴
RUN apk add --no-cache git ca-certificates tzdata

# 設置工作目錄
WORKDIR /app

# 複製 go mod 文件
COPY go.mod go.sum ./

# 下載依賴
RUN go mod download

# 複製源碼
COPY . .

# 構建應用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/go-server

# 運行階段
FROM alpine:latest

# 安裝 ca-certificates 和 tzdata
RUN apk --no-cache add ca-certificates tzdata

# 創建非 root 用戶
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# 設置工作目錄
WORKDIR /root/

# 從構建階段複製二進制文件
COPY --from=builder /app/main .

# 設置權限
RUN chown -R appuser:appgroup /root/
USER appuser

# 暴露端口
EXPOSE 8080

# 健康檢查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# 運行應用
CMD ["./main"]
