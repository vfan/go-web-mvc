import axios from 'axios';

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
    // 根据约定，处理后端返回的数据格式
    const res = response.data;
    if (res.code !== 0) {
      // 处理后端业务逻辑错误
      console.error(res.msg || '请求出错');
      return Promise.reject(new Error(res.msg || '请求出错'));
    } else {
      return res.data;
    }
  },
  error => {
    console.error('网络请求失败', error);
    return Promise.reject(error);
  }
);

export default api; 