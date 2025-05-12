import { useState } from 'react';
import { Form, Input, Button, Card, message, Typography } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';

const { Title } = Typography;

// 登录表单接口定义
interface LoginFormValues {
  username: string;
  password: string;
}

// 登录页面元数据
export function meta() {
  return [
    { title: "登录 - Go Web MVC应用" },
    { name: "description", content: "用户登录页面" },
  ];
}

// 登录页面组件
export default function Login() {
  const [loading, setLoading] = useState(false);

  // 处理登录提交
  const handleSubmit = async (values: LoginFormValues) => {
    setLoading(true);
    
    try {
      // 这里应该调用真实的登录API
      // 模拟API调用
      await new Promise(resolve => setTimeout(resolve, 1000));
      
      if (values.username === 'admin' && values.password === 'admin') {
        message.success('登录成功');
        // 登录成功后重定向到首页
        window.location.href = '/home';
      } else {
        message.error('用户名或密码错误');
      }
    } catch (error) {
      message.error('登录失败，请稍后重试');
      console.error('登录错误:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <Card className="w-96 shadow-md">
        <div className="text-center mb-6">
          <Title level={3}>Go Web MVC</Title>
          <p className="text-gray-500">请登录您的账号</p>
        </div>
        
        <Form
          name="login"
          initialValues={{ remember: true }}
          onFinish={handleSubmit}
          layout="vertical"
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
              className="w-full" 
              size="large"
              loading={loading}
            >
              登录
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
} 