import axios from 'axios';

// API响应的通用接口
export interface ApiResponse<T> {
  code: number;
  msg: string;
  data: T;
}

// 创建 axios 实例
const api = axios.create({
  baseURL: '/api', // 后端 API 的基础 URL
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  }
});

// 请求拦截器
api.interceptors.request.use(
  config => {
    // 可以在这里处理 token 等认证信息
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// 响应拦截器
api.interceptors.response.use(
  response => {
    // 直接返回原始响应，让各个API服务自己处理数据提取
    return response;
  },
  error => {
    // 处理HTTP错误（网络错误或服务器错误）
    if (error.response && error.response.status === 401) {
      // 处理未授权情况，如重定向到登录页面
      localStorage.removeItem('token');
      window.location.href = '/login';
    }
    
    // 显示错误消息
    const errorMsg = error.response?.data?.msg || '网络请求失败';
    console.error(errorMsg, error);
    
    return Promise.reject(error);
  }
);

export default api; 