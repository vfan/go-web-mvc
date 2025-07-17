import { createHashRouter, RouterProvider, Navigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import type { ReactNode } from 'react';

import App from '../App';
import Home from '../pages/Home';
import Login from '../pages/Login';
import NotFound from '../pages/NotFound';
import UserList from '../pages/User';
import UniversityList from '../pages/University';
import api from '../utils/api';
import RouterDemo from '../pages/RouterDemo/Demo';



// 受保护的路由组件
interface ProtectedRouteProps {
  children: ReactNode;
}

const ProtectedRoute = ({ children }: ProtectedRouteProps) => {
  const [isChecking, setIsChecking] = useState(true);
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const checkAuth = async () => {
      try {
        // 尝试访问一个需要认证的API端点来验证登录状态
        // 这里假设有一个 /api/auth/me 端点可以检查当前用户
        await api.get('/auth/me');
        setIsLoggedIn(true);
      } catch {
        setIsLoggedIn(false);
      } finally {
        setIsChecking(false);
      }
    };

    checkAuth();
  }, []);

  // 检查中时显示简单的加载提示
  if (isChecking) {
    return <div>正在检查登录状态...</div>;
  }

  // 如果没有登录，重定向到登录页面
  if (!isLoggedIn) {
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
const router = createHashRouter([
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
        path: 'university',
        element: <UniversityList />,
      },
      {
        path: 'settings',
        element: <Settings />,
      },
      {
        path: 'router',
        element: <RouterDemo />,
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