import { useEffect } from 'react';
import { Form, Input, Button, Select, message } from 'antd';
import type { User } from '../../+types/home';
import { createUser, updateUser } from '../../../services/userService';

const { Option } = Select;

// 用户表单属性接口
interface UserFormProps {
  user: User | null;
  onSuccess: () => void;
  onCancel: () => void;
}

export default function UserForm({ user, onSuccess, onCancel }: UserFormProps) {
  const [form] = Form.useForm();
  const [messageApi, contextHolder] = message.useMessage();
  const isEdit = !!user;

  // 表单初始化
  useEffect(() => {
    if (user) {
      form.setFieldsValue({
        email: user.email,
        role: user.role,
        password: '', // 编辑时不显示密码，但保留字段用于可选修改
      });
    } else {
      form.resetFields();
    }
  }, [user, form]);

  // 处理表单提交
  const handleSubmit = async (values: any) => {
    try {
      if (isEdit && user) {
        // 如果没有输入密码，移除密码字段
        if (!values.password) {
          delete values.password;
        }
        
        await updateUser(user.id, values);
        messageApi.success('更新用户成功');
      } else {
        await createUser(values);
        messageApi.success('创建用户成功');
      }
      
      onSuccess();
    } catch (error) {
      const action = isEdit ? '更新' : '创建';
      messageApi.error(`${action}用户失败`);
      console.error(`${action}用户错误:`, error);
    }
  };

  return (
    <Form
      form={form}
      layout="vertical"
      onFinish={handleSubmit}
    >
      {contextHolder}

      <Form.Item
        name="email"
        label="邮箱"
        rules={[
          { required: true, message: '请输入邮箱' },
          { type: 'email', message: '请输入有效的邮箱地址' }
        ]}
      >
        <Input placeholder="请输入邮箱" />
      </Form.Item>

      <Form.Item
        name="password"
        label="密码"
        rules={[
          { required: !isEdit, message: '请输入密码' },
          { min: 6, message: '密码长度不能少于6个字符' }
        ]}
        extra={isEdit ? "留空表示不修改密码" : ""}
      >
        <Input.Password placeholder={isEdit ? "留空表示不修改" : "请输入密码"} />
      </Form.Item>

      <Form.Item
        name="role"
        label="角色"
        rules={[{ required: true, message: '请选择角色' }]}
      >
        <Select placeholder="请选择角色">
          <Option value="admin">管理员</Option>
          <Option value="user">普通用户</Option>
        </Select>
      </Form.Item>

      <Form.Item className="flex justify-end gap-2 mb-0 mt-4">
        <Button onClick={onCancel}>取消</Button>
        <Button type="primary" htmlType="submit">
          {isEdit ? '更新' : '创建'}
        </Button>
      </Form.Item>
    </Form>
  );
} 