import { type RouteConfig, index, route } from "@react-router/dev/routes";

// 定义路由配置
export default [
  // 将登录页面作为默认路由
  index("routes/login.tsx"),
  // 添加首页路由
  route("home", "routes/home.tsx"),
  // 添加用户管理相关路由
  route("users", "routes/users/index.tsx"),
  route("users/:id", "routes/users/[id].tsx")
] satisfies RouteConfig;
