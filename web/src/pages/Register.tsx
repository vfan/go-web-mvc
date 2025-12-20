import { useState } from 'react';
import { Form, Input, Button, message, Card, Select } from 'antd';
import { UserOutlined, LockOutlined, MailOutlined } from '@ant-design/icons';
import { useNavigate, Link } from 'react-router-dom';
import api from '../utils/api';
import type { ApiResponse } from '../utils/api';

interface RegisterParams {
  email: string;
  username: string;
  password: string;
  role: number;
}

function Register() {
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();
  const [form] = Form.useForm();

  const onFinish = async (values: RegisterParams) => {
    try {
      setLoading(true);
      // 调用注册接口
      const response = await api.post<ApiResponse<any>>('/auth/register', values);
      
      if (response.data.code === 0) {
        message.success('注册成功，请登录');
        navigate('/login');
      } else {
        message.error(response.data.msg || '注册失败');
      }
    } catch (error) {
      console.error('注册失败:', error);
      message.error('注册失败，请稍后重试');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gradient-to-r from-blue-50 to-indigo-100">
      <div className="w-full max-w-md mx-auto">
        <div className="text-center mb-6">
          <h1 className="text-3xl font-bold text-indigo-700">Go Web MVC 系统</h1>
          <p className="text-gray-600 mt-2">注册新账号</p>
        </div>
        
        <Card className="w-full shadow-lg rounded-lg border-0" style={{ maxWidth: '400px', margin: '0 auto' }}>
          <h2 className="text-center text-2xl font-semibold mb-6 text-gray-700">用户注册</h2>
          <Form
            form={form}
            name="register"
            initialValues={{ role: 2 }}
            onFinish={onFinish}
            layout="vertical"
            size="large"
            style={{ maxWidth: '360px', margin: '0 auto' }}
          >
            <Form.Item
              name="username"
              rules={[{ required: true, message: '请输入用户名' }]}
            >
              <Input 
                prefix={<UserOutlined className="text-gray-400" />} 
                placeholder="用户名"
                className="rounded-md"
              />
            </Form.Item>
            
            <Form.Item
              name="email"
              rules={[
                { required: true, message: '请输入邮箱' },
                { type: 'email', message: '请输入有效的邮箱地址' }
              ]}
            >
              <Input 
                prefix={<MailOutlined className="text-gray-400" />} 
                placeholder="邮箱"
                className="rounded-md"
              />
            </Form.Item>
            
            <Form.Item
              name="password"
              rules={[
                { required: true, message: '请输入密码' },
                { min: 6, message: '密码至少6位' }
              ]}
            >
              <Input.Password 
                prefix={<LockOutlined className="text-gray-400" />} 
                placeholder="密码"
                className="rounded-md"
              />
            </Form.Item>
            
            <Form.Item
              name="role"
              label="角色"
              rules={[{ required: true, message: '请选择角色' }]}
            >
              <Select placeholder="请选择角色">
                <Select.Option value={2}>普通用户</Select.Option>
                <Select.Option value={1}>管理员</Select.Option>
              </Select>
            </Form.Item>

            <Form.Item className="mt-6">
              <Button 
                type="primary" 
                htmlType="submit" 
                loading={loading} 
                block
                className="h-10 bg-indigo-600 hover:bg-indigo-700 border-0 rounded-md"
              >
                注册
              </Button>
            </Form.Item>
            
            <div className="text-center">
              <Link to="/login" className="text-indigo-600 hover:text-indigo-800">
                已有账号？立即登录
              </Link>
            </div>
          </Form>
        </Card>
      </div>
    </div>
  );
}

export default Register;






