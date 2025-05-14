# Docker 使用指南

本文档提供了如何使用 Docker 构建和运行服务端应用的详细说明。

## 先决条件

- Docker 安装 (推荐 20.10.0 或更高版本)
- Docker Compose 安装 (推荐 1.29.0 或更高版本)

## 单独使用 Dockerfile 构建和运行

### 构建镜像

在项目根目录（包含 Dockerfile 的目录）执行以下命令：

```bash
docker build -t mvc-demo:latest .
```

### 运行容器

```bash
docker run -d -p 8080:8080 \
  --name mvc-demo \
  -e DB_HOST=your_db_host \
  -e DB_PORT=3306 \
  -e DB_USER=root \
  -e DB_PASSWORD=your_password \
  -e DB_NAME=mvc_demo \
  -e GIN_MODE=release \
  mvc-demo:latest
```

## 使用 Docker Compose

使用 Docker Compose 可以同时启动应用和数据库，并自动配置网络连接。

### 启动服务

```bash
docker-compose up -d
```

这将启动应用服务和 MySQL 数据库。数据库将使用 `./db/db.sql` 初始化。

### 停止服务

```bash
docker-compose down
```

若要同时删除卷，添加 `-v` 参数：

```bash
docker-compose down -v
```

## 环境变量配置

应用支持以下环境变量：

| 变量名 | 说明 | 默认值 |
|-------|------|-------|
| SERVER_PORT | 服务端口 | :8080 |
| GIN_MODE | Gin框架运行模式 (debug/release) | debug |
| DB_HOST | 数据库主机 | localhost |
| DB_PORT | 数据库端口 | 3306 |
| DB_USER | 数据库用户名 | root |
| DB_PASSWORD | 数据库密码 | - |
| DB_NAME | 数据库名称 | mvc_demo |
| DB_CHARSET | 数据库字符集 | utf8mb4 |
| JWT_SECRET | JWT密钥 | - |
| JWT_EXPIRES | JWT过期时间 | 24h |

## CI/CD 集成

在 CI/CD 流程中，可以使用以下命令构建和推送 Docker 镜像到镜像仓库：

```bash
# 登录到镜像仓库
docker login $REGISTRY_URL -u $REGISTRY_USER -p $REGISTRY_PASSWORD

# 构建镜像
docker build -t $REGISTRY_URL/mvc-demo:$VERSION .

# 推送镜像
docker push $REGISTRY_URL/mvc-demo:$VERSION
```

## 生产环境注意事项

1. 在生产环境中，建议将敏感配置（如数据库密码）通过环境变量或安全的配置管理工具提供，而不是硬编码在配置文件中。
2. 设置 GIN_MODE=release 以禁用调试功能并提高性能。
3. 考虑使用外部数据库服务而不是容器化的数据库，以确保数据持久性和备份。
4. 实现健康检查和监控，确保服务可用性。 