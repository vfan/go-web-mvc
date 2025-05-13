import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
// 导入路由
import Router from './router'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Router />
  </StrictMode>,
)
