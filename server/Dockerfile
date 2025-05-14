# 第一阶段：构建应用
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装必要的系统依赖
RUN apk add --no-cache gcc musl-dev

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载所有依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# 第二阶段：创建最小运行镜像
FROM alpine:latest

# 安装必要的运行时依赖
RUN apk --no-cache add ca-certificates tzdata

# 设置中国时区
ENV TZ=Asia/Shanghai

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译好的应用
COPY --from=builder /app/app .

# 复制配置文件（如果有）
COPY --from=builder /app/.env* ./

# 暴露应用端口
EXPOSE 8080

# 运行应用
CMD ["./app"]

# 提供健康检查
HEALTHCHECK --interval=30s --timeout=3s --retries=3 CMD wget --spider -q http://localhost:8080/health || exit 1 