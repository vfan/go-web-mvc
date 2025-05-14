# 使用 docker build -t go-web-mvc-frontend . 构建镜像

# 第一阶段：构建应用
FROM node:20-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制package.json和package-lock.json
COPY package*.json ./

# 安装依赖
RUN npm install

# 复制所有文件
COPY . .

# 构建应用
RUN npm run build

# 第二阶段：使用Nginx托管
FROM nginx:alpine

# 将构建产物从builder阶段复制到Nginx目录
COPY --from=builder /app/dist /usr/share/nginx/html

# 复制自定义的Nginx配置
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 暴露80端口
EXPOSE 80

# 启动Nginx
CMD ["nginx", "-g", "daemon off;"] 