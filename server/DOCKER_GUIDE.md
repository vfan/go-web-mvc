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

下面示例的DB_HOST=mysql,表示数据库在docker中启动，并且容器名称为mysql。

```bash
docker run -d -p 8080:8080 \
  --name mvc-demo \
  -e DB_HOST=mysql \
  -e DB_PORT=3306 \
  -e DB_USER=web \
  -e DB_PASSWORD=golang@2025 \
  -e DB_NAME=student_management \
  -e GIN_MODE=release \
  -e JWT_SECRET_KEY=your-secret-key-change-in-production \
  -e JWT_TOKEN_EXPIRY=24 \
  -e JWT_REFRESH_EXPIRY=168 \
  -e JWT_ISSUER=student-management-system \
  mvc-demo:latest
```


1. 如果数据库是在本地以docker启动，可以配置network,让其它容器连接到数据库的网络，从而连接到数据库。

```bash
docker network create mvc-network
docker run -d --name mysql --network mvc-network -e MYSQL_ROOT_PASSWORD=123456 mysql:8.0

# 或者把已经启动的容器加入到这个网络,这样在go服务中，可以使用mysql容器名直接连接数据库
docker network connect mvc-network mysql
docker network connect mvc-network mvc-demo

 

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

