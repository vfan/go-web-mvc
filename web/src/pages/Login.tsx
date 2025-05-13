import { useState } from 'react';
import { Form, Input, Button, message, Card } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import api from '../utils/api';
import type { ApiResponse } from '../utils/api';

interface LoginResponse {
  token: string;
  token_type: string;
  expires_in: number;
}

interface LoginParams {
  email: string;
  password: string;
}

function Login() {
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const onFinish = async (values: LoginParams) => {
    try {
      setLoading(true);
      // 调用登录接口
      const response = await api.post<ApiResponse<LoginResponse>>('/auth/login', values);
      
      if (response.data.code === 0) {
        // 登录成功后存储令牌
        const loginData = response.data.data;
        localStorage.setItem('token', loginData.token);
        localStorage.setItem('userInfo', JSON.stringify({
          email: values.email
        }));
        
        message.success('登录成功');
        navigate('/');
      } else {
        message.error(response.data.msg || '登录失败');
      }
    } catch (error) {
      console.error('登录失败:', error);
      message.error('登录失败，请检查邮箱和密码');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gradient-to-r from-blue-50 to-indigo-100">
      <div className="w-full max-w-xs mx-auto">
        <div className="text-center mb-6">
          <h1 className="text-3xl font-bold text-indigo-700">Go Web MVC 系统</h1>
          <p className="text-gray-600 mt-2">请登录以继续使用</p>
        </div>
        
        <Card className="w-full shadow-lg rounded-lg border-0" style={{ maxWidth: '320px', margin: '0 auto' }}>
          <h2 className="text-center text-2xl font-semibold mb-6 text-gray-700">用户登录</h2>
          <Form
            name="login"
            initialValues={{ remember: true }}
            onFinish={onFinish}
            layout="vertical"
            size="large"
            style={{ maxWidth: '280px', margin: '0 auto' }}
          >
            <Form.Item
              name="email"
              rules={[
                { required: true, message: '请输入邮箱' },
                { type: 'email', message: '请输入有效的邮箱地址' }
              ]}
            >
              <Input 
                prefix={<UserOutlined className="text-gray-400" />} 
                placeholder="邮箱"
                className="rounded-md"
                style={{ width: '100%', maxWidth: '280px' }}
              />
            </Form.Item>
            <Form.Item
              name="password"
              rules={[{ required: true, message: '请输入密码' }]}
            >
              <Input.Password 
                prefix={<LockOutlined className="text-gray-400" />} 
                placeholder="密码"
                className="rounded-md"
                style={{ width: '100%', maxWidth: '280px' }}
              />
            </Form.Item>
            <Form.Item className="mt-6">
              <Button 
                type="primary" 
                htmlType="submit" 
                loading={loading} 
                block
                className="h-10 bg-indigo-600 hover:bg-indigo-700 border-0 rounded-md"
              >
                登录
              </Button>
            </Form.Item>
          </Form>
        </Card>
        
        <div className="text-center mt-6 text-gray-500 text-sm">
          © {new Date().getFullYear()} Go Web MVC 系统 版权所有
        </div>
      </div>
    </div>
  );
}

export default Login; 