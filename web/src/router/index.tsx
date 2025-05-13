import { createBrowserRouter, RouterProvider, Navigate } from 'react-router-dom';
import App from '../App';
import Home from '../pages/Home';
import Login from '../pages/Login';
import NotFound from '../pages/NotFound';
import UserList from '../pages/User';
import type { ReactNode } from 'react';

// 检查是否已登录
const isAuthenticated = () => {
  return !!localStorage.getItem('token');
};

// 受保护的路由组件
interface ProtectedRouteProps {
  children: ReactNode;
}

const ProtectedRoute = ({ children }: ProtectedRouteProps) => {
  if (!isAuthenticated()) {
    return <Navigate to="/login" />;
  }
  return <>{children}</>;
};

// 设置页面组件（暂时使用占位组件）
const Settings = () => (
  <div className="p-4">
    <h2 className="text-2xl font-bold mb-4">系统设置</h2>
    <p>系统设置页面，正在开发中...</p>
  </div>
);

// 路由配置
const router = createBrowserRouter([
  // 登录页面（不使用主布局）
  {
    path: '/login',
    element: <Login />,
  },
  // 使用App布局的页面（需要登录）
  {
    path: '/',
    element: (
      <ProtectedRoute>
        <App />
      </ProtectedRoute>
    ),
    children: [
      {
        index: true,
        element: <Home />,
      },
      {
        path: 'user',
        element: <UserList />,
      },
      {
        path: 'settings',
        element: <Settings />,
      },
      {
        path: '*',
        element: <NotFound />,
      },
    ],
  },
]);

export default function Router() {
  return <RouterProvider router={router} />;
} 