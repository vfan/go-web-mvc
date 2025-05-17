import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
// 导入路由
import Router from './router'
// 导入Ant Design的App组件提供全局配置
import { message, App as AntdApp } from 'antd'

// 全局配置message
message.config({
  top: 100,
  duration: 2,
  maxCount: 3,
});

createRoot(document.getElementById('root')!).render(
    <AntdApp>
      <Router />
    </AntdApp>
)
