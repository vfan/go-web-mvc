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

// 处理未登录状态，跳转到登录页面
const handleUnauthorized = () => {
  // 不再需要清除localStorage中的token
  localStorage.removeItem('userInfo');
  
  // 如果不在登录页，则跳转
  if (window.location.pathname !== '/login') {
    window.location.href = '/login';
  }
};

// 请求拦截器
api.interceptors.request.use(
  config => {
    // 不再需要从localStorage获取token并添加到header中
    // 因为cookie会自动随请求发送
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// 响应拦截器
api.interceptors.response.use(
  response => {
    // 检查业务逻辑错误码
    if (response.data && response.data.code === -2) {
      // code为-2表示未登录或token无效
      handleUnauthorized();
      return Promise.reject(new Error(response.data.msg || '未登录，请先登录'));
    }
    // 直接返回原始响应，让各个API服务自己处理数据提取
    return response;
  },
  error => {
    // 处理HTTP错误（网络错误或服务器错误）
    if (error.response && error.response.status === 401) {
      // 处理未授权情况，如重定向到登录页面
      handleUnauthorized();
    }
    
    // 显示错误消息
    const errorMsg = error.response?.data?.msg || '网络请求失败';
    console.error(errorMsg, error);
    
    return Promise.reject(error);
  }
);

export default api; 