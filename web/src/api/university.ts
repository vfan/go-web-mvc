import api from '../utils/api';
import type { ApiResponse } from '../utils/api';

export interface University {
  id: number;
  name: string;
  created_at: string;
  updated_at: string;
  created_by: number | null;
  updated_by: number | null;
  deleted_at?: {
    Time: string;
    Valid: boolean;
  } | null;
}

export interface UniversityListResponse {
  total: number;
  list: University[];
}

export interface UniversityCreateParams {
  name: string;
}

export interface UniversityUpdateParams {
  name: string;
}

// 检查API响应状态
const checkResponse = <T>(response: ApiResponse<T>): T => {
  // 不处理已在api.ts中处理的code=-2未登录情况
  if (response.code !== 0) {
    // 直接返回错误信息，让UI层处理显示
    throw new Error(response.msg || '请求失败');
  }
  return response.data;
};

// 获取大学列表（分页）
export const getUniversityList = async (page = 1, pageSize = 10, showDeleted = false) => {
  try {
    const response = await api.get<ApiResponse<UniversityListResponse>>('/universities', {
      params: { page, page_size: pageSize, show_deleted: showDeleted },
    });
    return checkResponse(response.data);
  } catch (error) {
    console.error('获取大学列表失败:', error);
    throw error;
  }
};

// 获取所有大学（不分页）
export const getAllUniversities = async () => {
  try {
    const response = await api.get<ApiResponse<University[]>>('/universities/all');
    return checkResponse(response.data);
  } catch (error) {
    console.error('获取所有大学失败:', error);
    throw error;
  }
};

// 获取大学详情
export const getUniversityDetail = async (id: number) => {
  try {
    const response = await api.get<ApiResponse<University>>(`/universities/${id}`);
    return checkResponse(response.data);
  } catch (error) {
    console.error('获取大学详情失败:', error);
    throw error;
  }
};

// 创建大学（管理员权限）
export const createUniversity = async (params: UniversityCreateParams) => {
  try {
    const response = await api.post<ApiResponse<University>>('/admin/universities', params);
    return checkResponse(response.data);
  } catch (error) {
    console.error('创建大学失败:', error);
    throw error;
  }
};

// 更新大学（管理员权限）
export const updateUniversity = async (id: number, params: UniversityUpdateParams) => {
  try {
    const response = await api.put<ApiResponse<University>>(`/admin/universities/${id}`, params);
    return checkResponse(response.data);
  } catch (error) {
    console.error('更新大学失败:', error);
    throw error;
  }
};

// 删除大学（管理员权限）
export const deleteUniversity = async (id: number) => {
  try {
    const response = await api.delete<ApiResponse<null>>(`/admin/universities/${id}`);
    return checkResponse(response.data);
  } catch (error) {
    console.error('删除大学失败:', error);
    throw error;
  }
};

// 恢复已删除的大学（管理员权限）
export const restoreUniversity = async (id: number) => {
  try {
    const response = await api.post<ApiResponse<null>>(`/admin/universities/${id}/restore`);
    return checkResponse(response.data);
  } catch (error) {
    console.error('恢复大学失败:', error);
    throw error;
  }
}; 