# Docker 使用指南

本文档说明如何**单独**构建与运行服务端镜像；前端 + 后端 + MySQL 的一体化编排见仓库根目录的 [`README.md`](../README.md) 与 [`docker-compose.yml`](../docker-compose.yml)。

## 先决条件

- Docker（推荐 20.10 及以上）
- Docker Compose（独立版或 Compose V2 插件均可）

## 单独使用 Dockerfile 构建和运行

服务端镜像的 `Dockerfile` 位于 `server/` 目录下，需在仓库根目录指定构建上下文：

### 构建镜像

```bash
docker build -t mvc-demo:latest -f server/Dockerfile ./server
```

或在 `server` 目录下执行：

```bash
cd server
docker build -t mvc-demo:latest .
```

### 运行容器（与 Docker 内的 MySQL 联调）

`DB_HOST=mysql` 表示通过 **容器名** 访问数据库，应用容器必须与 MySQL 在同一自定义网络上，并使用与 `MYSQL_ROOT_PASSWORD` 一致的 **root** 密码。

在**仓库根目录**执行（以便挂载 `./server/db/db.sql`）：

```bash
docker network create mvc-network

docker run -d --name mysql --network mvc-network \
  -e MYSQL_ROOT_PASSWORD=123456 \
  -e MYSQL_DATABASE=student_management \
  -v "$(pwd)/server/db/db.sql:/docker-entrypoint-initdb.d/db.sql:ro" \
  mysql:8.0
```

首次启动时请等待数十秒，待 MySQL 完成初始化后再启动后端：

```bash
docker run -d -p 8080:8080 --name mvc-demo --network mvc-network \
  -e DB_HOST=mysql \
  -e DB_PORT=3306 \
  -e DB_USER=root \
  -e DB_PASSWORD=123456 \
  -e DB_NAME=student_management \
  -e GIN_MODE=release \
  -e JWT_SECRET_KEY=your-secret-key-change-in-production \
  -e JWT_TOKEN_EXPIRY=24 \
  -e JWT_REFRESH_EXPIRY=168 \
  -e JWT_ISSUER=student-management-system \
  mvc-demo:latest
```

若 MySQL 或后端容器已存在且未加入同一网络，可使用：

```bash
docker network connect mvc-network mysql
docker network connect mvc-network mvc-demo
```

### 构建镜像的目标平台（amd64）

`server/Dockerfile` 构建阶段设置了 `GOOS=linux`、`GOARCH=amd64`，镜像可在常见 x86 云主机上运行。若在 **ARM64 Linux 服务器**上部署，请将构建参数改为 `GOARCH=arm64`（或直接使用 `docker buildx build --platform linux/arm64`），否则二进制无法在目标 CPU 上运行。

开发机为 Apple 芯片时，在本机用 Docker 构建 amd64 镜像一般会通过 QEMU 模拟完成，通常无需改 Dockerfile。

## 使用 Docker Compose

使用 Docker Compose 可以同时启动应用和数据库，并自动配置网络连接。

### 启动服务

在**仓库根目录**（与 `docker-compose.yml` 同级）执行：

```bash
docker-compose up -d
```

也可使用 Compose V2：`docker compose up -d`。

这将启动前端（Nginx）、后端与 MySQL。首次启动时，MySQL 会使用 `./server/db/db.sql` 初始化数据库（见 `docker-compose.yml` 中的挂载配置）。

### 查看日志与重建

```bash
docker compose logs -f server
docker compose up -d --build
```

### 停止服务

```bash
docker-compose down
```

若要同时删除卷，添加 `-v` 参数：

```bash
docker-compose down -v
```

## 环境变量配置

应用支持的主要环境变量如下（与 `server/config`、`server/.env.example` 一致）：

| 变量名 | 说明 | 默认值 |
|-------|------|-------|
| SERVER_PORT | 服务监听地址 | :8080 |
| GIN_MODE | Gin 运行模式 (debug/release) | debug |
| DB_HOST | 数据库主机 | localhost |
| DB_PORT | 数据库端口 | 3306 |
| DB_USER | 数据库用户名 | root |
| DB_PASSWORD | 数据库密码 | （空） |
| DB_NAME | 数据库名称 | student_management（与 `server/db/db.sql` 一致；部署时建议显式设置） |
| DB_CHARSET | 字符集 | utf8mb4 |
| JWT_SECRET_KEY | JWT 签名密钥 | your-secret-key-change-in-production |
| JWT_TOKEN_EXPIRY | 访问令牌有效期（小时） | 24 |
| JWT_REFRESH_EXPIRY | 刷新令牌有效期（小时） | 168 |
| JWT_ISSUER | JWT 签发者 | student-management-system |

