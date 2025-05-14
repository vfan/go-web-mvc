# 学生管理系统 API 文档

## 目录

- [概述](#概述)
- [认证机制](#认证机制)
- [响应格式](#响应格式)
- [API接口](#api接口)
  - [认证接口](#认证接口)
  - [用户管理](#用户管理)
  - [大学管理](#大学管理)
  - [学生管理](#学生管理)

## 概述

本文档详细描述了学生管理系统后端API接口，包括认证、用户管理、大学信息管理和学生信息管理等功能。

## 认证机制

系统使用JWT（JSON Web Token）进行认证。除登录接口外，所有接口都需要在请求头中携带token。

### 认证流程

1. 调用登录接口获取token
2. 在后续请求的Header中添加：`Authorization: Bearer {token}`

### 权限级别

系统有两种角色：
- **管理员(role=1)**: 可以进行所有操作，包括增删改查
- **普通用户(role=2)**: 只能进行查询操作

## 响应格式

所有API返回统一的JSON格式：

```json
{
  "code": 0,         // 状态码，0表示成功，小于0表示错误
  "msg": "成功",      // 响应消息
  "data": {}         // 响应数据，可能是对象、数组或null
}
```

常见状态码：
- 0: 请求成功
- -1: 请求参数错误
- -2: 未授权（未登录或token无效）
- -3: 权限不足
- -4: 资源不存在
- -5: 服务器内部错误

## API接口

### 认证接口

#### 登录

- **URL**: `/api/auth/login`
- **方法**: POST
- **权限**: 无需认证
- **请求参数**:

  ```json
  {
    "email": "user@example.com",
    "password": "yourpassword"
  }
  ```

- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "登录成功",
    "data": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "token_type": "Bearer",
      "expires_in": 86400
    }
  }
  ```

### 用户管理

#### 获取用户列表

- **URL**: `/api/users`
- **方法**: GET
- **权限**: 需认证（所有用户）
- **查询参数**:
  - `page`: 页码，默认1
  - `page_size`: 每页条数，默认10
- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "成功",
    "data": {
      "total": 100,
      "list": [
        {
          "id": 1,
          "email": "admin@example.com",
          "role": 1,
          "last_login_time": "2023-05-10T15:30:45Z",
          "status": 1,
          "created_at": "2023-05-01T10:00:00Z",
          "updated_at": "2023-05-10T15:30:45Z"
        },
        // ...更多用户
      ],
      "page": 1,
      "size": 10
    }
  }
  ```

#### 获取用户详情

- **URL**: `/api/users/{id}`
- **方法**: GET
- **权限**: 需认证（所有用户）
- **路径参数**:
  - `id`: 用户ID
- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "成功",
    "data": {
      "id": 1,
      "email": "admin@example.com",
      "role": 1,
      "last_login_time": "2023-05-10T15:30:45Z",
      "status": 1,
      "created_at": "2023-05-01T10:00:00Z",
      "updated_at": "2023-05-10T15:30:45Z"
    }
  }
  ```

#### 创建用户

- **URL**: `/api/admin/users`
- **方法**: POST
- **权限**: 需认证（仅管理员）
- **请求参数**:

  ```json
  {
    "email": "newuser@example.com",
    "password": "userpassword",
    "role": 2,
    "status": 1
  }
  ```

- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "创建成功",
    "data": {
      "id": 2,
      "email": "newuser@example.com",
      "role": 2,
      "status": 1,
      "created_at": "2023-05-11T09:30:00Z",
      "updated_at": "2023-05-11T09:30:00Z"
    }
  }
  ```

#### 更新用户

- **URL**: `/api/admin/users/{id}`
- **方法**: PUT
- **权限**: 需认证（仅管理员）
- **路径参数**:
  - `id`: 用户ID
- **请求参数**:

  ```json
  {
    "email": "updated@example.com",
    "password": "newpassword",  // 可选，不修改密码时可不传
    "role": 2,
    "status": 1
  }
  ```

- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "更新成功",
    "data": {
      "id": 2,
      "email": "updated@example.com",
      "role": 2,
      "status": 1,
      "created_at": "2023-05-11T09:30:00Z",
      "updated_at": "2023-05-11T10:15:00Z"
    }
  }
  ```

#### 删除用户

- **URL**: `/api/admin/users/{id}`
- **方法**: DELETE
- **权限**: 需认证（仅管理员）
- **路径参数**:
  - `id`: 用户ID
- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "删除成功",
    "data": null
  }
  ```

### 大学管理

#### 获取大学列表

- **URL**: `/api/universities`
- **方法**: GET
- **权限**: 需认证（所有用户）
- **查询参数**:
  - `page`: 页码，默认1
  - `page_size`: 每页条数，默认10
- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "成功",
    "data": {
      "total": 50,
      "list": [
        {
          "id": 1,
          "name": "北京大学",
          "created_at": "2023-05-01T10:00:00Z",
          "updated_at": "2023-05-01T10:00:00Z",
          "created_by": 1,
          "updated_by": 1
        },
        // ...更多大学
      ]
    }
  }
  ```

#### 获取所有大学（不分页）

- **URL**: `/api/universities/all`
- **方法**: GET
- **权限**: 需认证（所有用户）
- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "成功",
    "data": [
      {
        "id": 1,
        "name": "北京大学",
        "created_at": "2023-05-01T10:00:00Z",
        "updated_at": "2023-05-01T10:00:00Z",
        "created_by": 1,
        "updated_by": 1
      },
      // ...所有大学
    ]
  }
  ```

#### 获取大学详情

- **URL**: `/api/universities/{id}`
- **方法**: GET
- **权限**: 需认证（所有用户）
- **路径参数**:
  - `id`: 大学ID
- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "成功",
    "data": {
      "id": 1,
      "name": "北京大学",
      "created_at": "2023-05-01T10:00:00Z",
      "updated_at": "2023-05-01T10:00:00Z",
      "created_by": 1,
      "updated_by": 1
    }
  }
  ```

#### 创建大学

- **URL**: `/api/admin/universities`
- **方法**: POST
- **权限**: 需认证（仅管理员）
- **请求参数**:

  ```json
  {
    "name": "清华大学"
  }
  ```

- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "创建成功",
    "data": {
      "id": 2,
      "name": "清华大学",
      "created_at": "2023-05-11T10:00:00Z",
      "updated_at": "2023-05-11T10:00:00Z",
      "created_by": 1,
      "updated_by": 1
    }
  }
  ```

#### 更新大学

- **URL**: `/api/admin/universities/{id}`
- **方法**: PUT
- **权限**: 需认证（仅管理员）
- **路径参数**:
  - `id`: 大学ID
- **请求参数**:

  ```json
  {
    "name": "清华大学(更新)"
  }
  ```

- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "更新成功",
    "data": {
      "id": 2,
      "name": "清华大学(更新)",
      "created_at": "2023-05-11T10:00:00Z",
      "updated_at": "2023-05-11T11:30:00Z",
      "created_by": 1,
      "updated_by": 1
    }
  }
  ```

#### 删除大学

- **URL**: `/api/admin/universities/{id}`
- **方法**: DELETE
- **权限**: 需认证（仅管理员）
- **路径参数**:
  - `id`: 大学ID
- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "删除成功",
    "data": null
  }
  ```

### 学生管理

#### 获取学生列表

- **URL**: `/api/students`
- **方法**: GET
- **权限**: 需认证（所有用户）
- **查询参数**:
  - `page`: 页码，默认1
  - `page_size`: 每页条数，默认10
  - `name`: 学生姓名（可选，模糊查询）
  - `university_id`: 大学ID（可选，精确查询）
  - `education`: 学历（可选，精确查询）
  - `graduation_year`: 毕业年份（可选，精确查询）
- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "成功",
    "data": {
      "total": 200,
      "list": [
        {
          "id": 1,
          "name": "张三",
          "email": "zhangsan@example.com",
          "gender": 1,
          "birthday": "2000-01-01",
          "phone": "13800138000",
          "university_id": 1,
          "university_name": "北京大学",
          "major": "计算机科学",
          "education": "本科",
          "graduation_year": 2023,
          "status": "在读",
          "avatar": "https://example.com/avatar/1.jpg",
          "created_at": "2023-05-01T10:00:00Z",
          "updated_at": "2023-05-01T10:00:00Z"
        },
        // ...更多学生
      ]
    }
  }
  ```

#### 获取学生详情

- **URL**: `/api/students/{id}`
- **方法**: GET
- **权限**: 需认证（所有用户）
- **路径参数**:
  - `id`: 学生ID
- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "成功",
    "data": {
      "id": 1,
      "name": "张三",
      "email": "zhangsan@example.com",
      "gender": 1,
      "birthday": "2000-01-01",
      "phone": "13800138000",
      "resume_path": "/uploads/resumes/1.pdf",
      "university_id": 1,
      "university_name": "北京大学",
      "major": "计算机科学",
      "education": "本科",
      "graduation_year": 2023,
      "status": "在读",
      "remarks": "优秀学生",
      "avatar": "https://example.com/avatar/1.jpg",
      "created_at": "2023-05-01T10:00:00Z",
      "updated_at": "2023-05-01T10:00:00Z",
      "created_by": 1,
      "updated_by": 1
    }
  }
  ```

#### 创建学生

- **URL**: `/api/admin/students`
- **方法**: POST
- **权限**: 需认证（仅管理员）
- **请求参数**:

  ```json
  {
    "name": "李四",
    "email": "lisi@example.com",
    "password": "studentpassword",
    "gender": 1,
    "birthday": "2001-02-03",
    "phone": "13900139000",
    "university_id": 2,
    "major": "软件工程",
    "education": "本科",
    "graduation_year": 2024,
    "status": "在读",
    "remarks": "转学生"
  }
  ```

- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "创建成功",
    "data": {
      "id": 2,
      "name": "李四",
      "email": "lisi@example.com",
      "gender": 1,
      "birthday": "2001-02-03",
      "phone": "13900139000",
      "university_id": 2,
      "university_name": "清华大学",
      "major": "软件工程",
      "education": "本科",
      "graduation_year": 2024,
      "status": "在读",
      "remarks": "转学生",
      "created_at": "2023-05-11T14:00:00Z",
      "updated_at": "2023-05-11T14:00:00Z",
      "created_by": 1,
      "updated_by": 1
    }
  }
  ```

#### 更新学生

- **URL**: `/api/admin/students/{id}`
- **方法**: PUT
- **权限**: 需认证（仅管理员）
- **路径参数**:
  - `id`: 学生ID
- **请求参数**:

  ```json
  {
    "name": "李四(已更新)",
    "email": "lisi_updated@example.com",
    "password": "newpassword",  // 可选，不修改密码时可不传
    "gender": 1,
    "birthday": "2001-02-03",
    "phone": "13900139001",
    "university_id": 2,
    "major": "人工智能",
    "education": "硕士",
    "graduation_year": 2025,
    "status": "在读",
    "remarks": "已转专业"
  }
  ```

- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "更新成功",
    "data": {
      "id": 2,
      "name": "李四(已更新)",
      "email": "lisi_updated@example.com",
      "gender": 1,
      "birthday": "2001-02-03",
      "phone": "13900139001",
      "university_id": 2,
      "university_name": "清华大学",
      "major": "人工智能",
      "education": "硕士",
      "graduation_year": 2025,
      "status": "在读",
      "remarks": "已转专业",
      "created_at": "2023-05-11T14:00:00Z",
      "updated_at": "2023-05-11T15:30:00Z",
      "created_by": 1,
      "updated_by": 1
    }
  }
  ```

#### 删除学生

- **URL**: `/api/admin/students/{id}`
- **方法**: DELETE
- **权限**: 需认证（仅管理员）
- **路径参数**:
  - `id`: 学生ID
- **响应示例**:

  ```json
  {
    "code": 0,
    "msg": "删除成功",
    "data": null
  }
  ```

## 错误响应示例

### 参数错误

```json
{
  "code": -1,
  "msg": "无效的请求参数",
  "data": null
}
```

### 认证失败

```json
{
  "code": -2,
  "msg": "未提供授权令牌",
  "data": null
}
```

### 权限不足

```json
{
  "code": -3,
  "msg": "需要管理员权限",
  "data": null
}
```

### 资源不存在

```json
{
  "code": -4,
  "msg": "未找到指定资源",
  "data": null
}
```

### 服务器错误

```json
{
  "code": -5,
  "msg": "服务器内部错误",
  "data": null
}
```
