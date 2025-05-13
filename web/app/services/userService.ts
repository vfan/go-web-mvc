// 用户管理相关API服务

import { request } from './api';
import type { ApiResponse, User, PaginationParams, PaginatedData } from '../routes/+types/home';

// API基础路径
const API_BASE = '/api/users';

/**
 * 获取用户列表
 * @param params 分页参数
 * @returns 分页后的用户列表
 */
export async function getUserList(params: PaginationParams): Promise<ApiResponse<PaginatedData<User>>> {
  return request<ApiResponse<PaginatedData<User>>>(`${API_BASE}?page=${params.page}&pageSize=${params.pageSize}`);
}

/**
 * 获取用户详情
 * @param id 用户ID
 * @returns 用户详情
 */
export async function getUserDetail(id: number): Promise<ApiResponse<User>> {
  return request<ApiResponse<User>>(`${API_BASE}/${id}`);
}

/**
 * 创建用户
 * @param user 用户数据
 * @returns 创建结果
 */
export async function createUser(user: Partial<User>): Promise<ApiResponse<User>> {
  return request<ApiResponse<User>>(API_BASE, {
    method: 'POST',
    body: JSON.stringify(user),
  });
}

/**
 * 更新用户
 * @param id 用户ID
 * @param user 用户数据
 * @returns 更新结果
 */
export async function updateUser(id: number, user: Partial<User>): Promise<ApiResponse<User>> {
  return request<ApiResponse<User>>(`${API_BASE}/${id}`, {
    method: 'PUT',
    body: JSON.stringify(user),
  });
}

/**
 * 删除用户
 * @param id 用户ID
 * @returns 删除结果
 */
export async function deleteUser(id: number): Promise<ApiResponse<null>> {
  return request<ApiResponse<null>>(`${API_BASE}/${id}`, {
    method: 'DELETE',
  });
} 