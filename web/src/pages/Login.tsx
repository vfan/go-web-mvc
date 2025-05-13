import { useState } from 'react';
import { Form, Input, Button, message, Card } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import api from '../utils/api';

function Login() {
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const onFinish = async (values: any) => {
    try {
      setLoading(true);
      // 调用登录接口
      await api.post('/user/login', values);
      
      // 登录成功后存储令牌
      localStorage.setItem('token', 'mock-token'); // 实际应该使用服务端返回的token
      localStorage.setItem('userInfo', JSON.stringify({
        username: values.username,
        role: 'user'
      }));
      
      message.success('登录成功');
      navigate('/');
    } catch (error) {
      console.error('登录失败:', error);
      message.error('登录失败，请检查用户名和密码');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="w-full max-w-md">
        <div className="text-center mb-8">
          <h1 className="text-3xl font-bold text-gray-800">Go Web MVC 系统</h1>
          <p className="text-gray-600 mt-2">请登录以继续使用</p>
        </div>
        
        <Card className="w-full shadow-lg rounded-lg">
          <h2 className="text-center text-2xl font-bold mb-6 text-gray-700">用户登录</h2>
          <Form
            name="login"
            initialValues={{ remember: true }}
            onFinish={onFinish}
            layout="vertical"
            size="large"
          >
            <Form.Item
              name="username"
              rules={[{ required: true, message: '请输入用户名' }]}
            >
              <Input 
                prefix={<UserOutlined className="text-gray-400" />} 
                placeholder="用户名" 
              />
            </Form.Item>
            <Form.Item
              name="password"
              rules={[{ required: true, message: '请输入密码' }]}
            >
              <Input.Password 
                prefix={<LockOutlined className="text-gray-400" />} 
                placeholder="密码" 
              />
            </Form.Item>
            <Form.Item>
              <Button 
                type="primary" 
                htmlType="submit" 
                loading={loading} 
                block
                className="h-10"
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