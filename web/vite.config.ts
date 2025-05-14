import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'


// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tailwindcss()],
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        // 如果后端 API 没有 /api 前缀，可以取消注释以下行来重写路径
        // rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
})
