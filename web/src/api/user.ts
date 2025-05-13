import api from '../utils/api';
import { message } from 'antd';
import type { ApiResponse } from '../utils/api';

export interface User {
  id: number;
  email: string;
  role: number;
  status: number;
  last_login_time?: string;
  created_at: string;
  updated_at: string;
}

export interface UserListResponse {
  total: number;
  items: User[];
}

export interface UserCreateParams {
  email: string;
  password: string;
  role: number;
  status: number;
}

export interface UserUpdateParams {
  email?: string;
  password?: string;
  role?: number;
  status?: number;
}

// 检查API响应状态
const checkResponse = <T>(response: ApiResponse<T>): T => {
  if (response.code !== 0) {
    // 显示错误消息
    message.error(response.msg || '请求失败');
    throw new Error(response.msg || '请求失败');
  }
  return response.data;
};

// 获取用户列表
export const getUserList = async (page = 1, pageSize = 10) => {
  try {
    const response = await api.get<ApiResponse<UserListResponse>>('/users', {
      params: { page, page_size: pageSize }
    });
    return checkResponse(response.data);
  } catch (error) {
    console.error('获取用户列表失败:', error);
    throw error;
  }
};

// 获取用户详情
export const getUserDetail = async (id: number) => {
  try {
    const response = await api.get<ApiResponse<User>>(`/users/${id}`);
    return checkResponse(response.data);
  } catch (error) {
    console.error('获取用户详情失败:', error);
    throw error;
  }
};

// 创建用户（管理员权限）
export const createUser = async (params: UserCreateParams) => {
  try {
    const response = await api.post<ApiResponse<User>>('/admin/users', params);
    return checkResponse(response.data);
  } catch (error) {
    console.error('创建用户失败:', error);
    throw error;
  }
};

// 更新用户（管理员权限）
export const updateUser = async (id: number, params: UserUpdateParams) => {
  try {
    const response = await api.put<ApiResponse<User>>(`/admin/users/${id}`, params);
    return checkResponse(response.data);
  } catch (error) {
    console.error('更新用户失败:', error);
    throw error;
  }
};

// 删除用户（管理员权限）
export const deleteUser = async (id: number) => {
  try {
    const response = await api.delete<ApiResponse<null>>(`/admin/users/${id}`);
    return checkResponse(response.data);
  } catch (error) {
    console.error('删除用户失败:', error);
    throw error;
  }
}; 