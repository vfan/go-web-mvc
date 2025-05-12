import { type RouteConfig, index, route } from "@react-router/dev/routes";

// 定义路由配置
export default [
  // 将登录页面作为默认路由
  index("routes/login.tsx"),
  // 添加首页路由
  route("home", "routes/home.tsx")
] satisfies RouteConfig;
