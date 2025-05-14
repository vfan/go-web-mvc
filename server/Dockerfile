# 第一阶段：构建Go应用
FROM golang:1.23-alpine AS builder

# 设置Go环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

# 复制go mod和sum文件，先复制go.mod和go.sum，再复制其他文件，是为了加快构建速度，因为go.mod和go.sum不会经常变化。原理是docker的缓存机制，如果go.mod和go.sum没有变化，就不会重新下载依赖。

COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN go build -o main .

# 第二阶段：运行阶段，alpine是轻量级的linux发行版，适合部署在docker中
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从builder阶段复制二进制文件
COPY --from=builder /app/main .

# 复制环境配置文件，也可以不需要，视具体情况
# COPY --from=builder /app/.env .


# 暴露8080端口
EXPOSE 8080

# 直接运行应用
CMD ["./main"] 