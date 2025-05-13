import React, { useState } from 'react';
import { Form, Input, Button, Card, message } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import './index.css';

// 登录表单数据类型
interface LoginForm {
  username: string;
  password: string;
}

const Login: React.FC = () => {
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  // 登录提交处理
  const handleSubmit = async (values: LoginForm) => {
    try {
      setLoading(true);
      
      // 这里应该调用真实的登录API，现在暂时模拟
      // 后续会实现实际的接口调用
      setTimeout(() => {
        // 模拟登录成功
        if (values.username === 'admin' && values.password === 'admin123') {
          // 存储用户信息到本地存储
          localStorage.setItem('token', 'mock-token');
          localStorage.setItem('userInfo', JSON.stringify({
            username: values.username,
            role: 'admin'
          }));
          
          message.success('登录成功');
          navigate('/dashboard');
        } else {
          message.error('用户名或密码错误');
        }
        setLoading(false);
      }, 1000);
    } catch (error) {
      setLoading(false);
      message.error('登录失败，请重试');
      console.error('登录错误:', error);
    }
  };

  return (
    <div className="login-container">
      <Card className="login-card" title="系统登录" bordered={false}>
        <Form
          name="login"
          className="login-form"
          initialValues={{ remember: true }}
          onFinish={handleSubmit}
        >
          <Form.Item
            name="username"
            rules={[{ required: true, message: '请输入用户名' }]}
          >
            <Input 
              prefix={<UserOutlined />} 
              placeholder="用户名" 
              size="large"
            />
          </Form.Item>
          <Form.Item
            name="password"
            rules={[{ required: true, message: '请输入密码' }]}
          >
            <Input.Password
              prefix={<LockOutlined />}
              placeholder="密码"
              size="large"
            />
          </Form.Item>

          <Form.Item>
            <Button 
              type="primary" 
              htmlType="submit" 
              className="login-button" 
              loading={loading}
              size="large"
              block
            >
              登录
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
};

export default Login; 