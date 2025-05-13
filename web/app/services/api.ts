// API服务工具类

/**
 * 封装fetch请求方法
 * @param url 请求地址
 * @param options 请求选项
 * @returns Promise
 */
export async function request<T>(url: string, options?: RequestInit): Promise<T> {
  // 获取存储的token
  const token = localStorage.getItem('token');
  const tokenType = localStorage.getItem('token_type') || 'Bearer';
  
  // 请求默认选项
  const defaultOptions: RequestInit = {
    headers: {
      'Content-Type': 'application/json',
    },
  };
  
  // 如果有token，添加到请求头
  if (token) {
    defaultOptions.headers = {
      ...defaultOptions.headers,
      'Authorization': `${tokenType} ${token}`,
    };
  }
  
  // 合并选项
  const fetchOptions = {
    ...defaultOptions,
    ...options,
    headers: {
      ...defaultOptions.headers,
      ...(options?.headers || {}),
    },
  };
  
  // 发起请求
  const response = await fetch(url, fetchOptions);
  
  // 如果响应码是401，可能是token过期，跳转到登录页
  if (response.status === 401) {
    localStorage.removeItem('token');
    window.location.href = '/';
    throw new Error('登录已过期，请重新登录');
  }
  
  // 解析响应结果
  const result = await response.json();
  
  // 检查业务状态码
  if (result.code !== 200) {
    throw new Error(result.message || '请求失败');
  }
  
  return result;
} 