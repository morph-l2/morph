# Build stage
FROM golang:1.21-alpine AS builder

# 安装必要的工具
RUN apk add --no-cache git make gcc musl-dev

WORKDIR /app

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建二进制文件
RUN make build

# Runtime stage
FROM alpine:latest

# 安装 ca-certificates
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建阶段复制二进制文件
COPY --from=builder /app/bin/gas-price-oracle .

# 创建日志目录
RUN mkdir -p /data/logs/morph-gas-oracle

# 暴露 metrics 端口
EXPOSE 6060

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:6060/health || exit 1

# 运行服务
ENTRYPOINT ["./gas-price-oracle"]

