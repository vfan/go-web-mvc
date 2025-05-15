# Go-Web-MVC 前端项目

本项目是一个基于React + TypeScript + Vite的前端应用，作为Go-Web-MVC全栈项目的前端部分。这是一个教学项目，旨在展示基于MVC架构的全栈Web应用开发。

## 项目架构

本前端项目使用以下技术栈：

- **React**: 用于构建用户界面的JavaScript库
- **TypeScript**: 添加静态类型系统，提高代码质量和开发体验
- **Vite**: 下一代前端构建工具，提供极速的开发服务器和优化的构建
- **Ant Design (antd)**: 企业级UI组件库
- **Tailwind CSS**: 实用优先的CSS框架
- **Axios**: 基于Promise的HTTP客户端，用于发送Ajax请求

## 目录结构

```
web/
├── public/              # 静态资源目录
├── src/                 # 源代码目录
│   ├── api/             # API请求相关
│   ├── assets/          # 静态资源
│   ├── components/      # 可复用组件
│   ├── hooks/           # 自定义React Hooks
│   ├── layouts/         # 布局组件
│   ├── pages/           # 页面组件
│   ├── store/           # 状态管理
│   ├── types/           # TypeScript类型定义
│   ├── utils/           # 工具函数
│   ├── App.tsx          # 应用入口组件
│   ├── main.tsx         # 应用入口文件
│   └── vite-env.d.ts    # Vite类型声明
├── .eslintrc.cjs        # ESLint配置
├── index.html           # HTML模板
├── package.json         # 项目依赖和脚本
├── tsconfig.json        # TypeScript配置
├── tailwind.config.js   # Tailwind CSS配置
└── vite.config.ts       # Vite配置
```

## 运行项目


### 安装依赖

```bash
cd web
npm install
# 或
yarn
```

### 开发环境运行

```bash
npm run dev
# 或
yarn dev
```

这将启动开发服务器，通常在 http://localhost:5173 可以访问。

### 构建生产版本

```bash
npm run build
# 或
yarn build
```

构建后的文件将位于 `dist` 目录中。

## 后端API交互

本项目前端通过Axios与后端API进行交互，后端API文档位于服务器端项目中的 `server/api.md`。

开发阶段，使用代理转发请求，在 `vite.config.ts` 中配置代理。因此开发阶段，需要启动后端服务。

```ts
server: {
  proxy: {
    '/api': 'http://localhost:8080'
  }
}
```

  - `code`: 0表示成功，小于0表示错误
  - `msg`: 错误信息
  - `data`: 返回的数据

## 开发指南

### 添加新页面

1. 在 `src/pages` 目录下创建新页面组件
2. 在路由配置中添加新路由
3. 根据需要添加相应的API请求函数

### 添加新组件

1. 在 `src/components` 目录下创建新组件
2. 组件应遵循项目的样式和代码规范
3. 为组件添加适当的TypeScript类型定义

## 学习资源

- [React文档](https://react.dev/)
- [TypeScript文档](https://www.typescriptlang.org/docs/)
- [Vite文档](https://vitejs.dev/guide/)
- [Ant Design文档](https://ant.design/docs/react/introduce-cn)
- [Tailwind CSS文档](https://tailwindcss.com/docs)
- [Axios文档](https://axios-http.com/docs/intro)
