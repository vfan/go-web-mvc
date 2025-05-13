import { createBrowserRouter, RouterProvider, Navigate } from 'react-router-dom';
import App from '../App';
import Home from '../pages/Home';
import Login from '../pages/Login';
import NotFound from '../pages/NotFound';
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
        path: '*',
        element: <NotFound />,
      },
    ],
  },
]);

export default function Router() {
  return <RouterProvider router={router} />;
} 