// 首页路由相关类型定义

export namespace Route {
  export interface MetaArgs {
    params: Record<string, string>;
    data: unknown;
  }
}

// 用户相关类型定义
export interface User {
  id: number;
  username: string;
  email: string;
  role: string;
  created_at: string;
  updated_at: string;
}

// API响应类型
export interface ApiResponse<T> {
  code: number;
  message: string;
  data?: T;
}

// 分页请求参数
export interface PaginationParams {
  page: number;
  pageSize: number;
}

// 分页响应数据
export interface PaginatedData<T> {
  list: T[];
  total: number;
  page: number;
  pageSize: number;
  totalPages: number;
} 