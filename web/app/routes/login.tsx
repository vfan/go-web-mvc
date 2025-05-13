import { useState } from 'react';
import { Form, Input, Button, Card, message, Typography, App } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';

const { Title } = Typography;

// 登录表单接口定义
interface LoginFormValues {
  email: string;
  password: string;
}

// 登录响应接口定义
interface LoginResponse {
  code: number;
  message: string;
  data?: {
    token: string;
    token_type: string;
    expires_in: number;
  };
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
  const [messageApi, contextHolder] = message.useMessage();

  // 处理登录提交
  const handleSubmit = async (values: LoginFormValues) => {
    setLoading(true);
    
    try {
      // 调用后端登录API
      const response = await fetch('/api/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(values),
      });
      
      const result: LoginResponse = await response.json();
      
      if (response.ok && result.code === 200) {
        // 登录成功，保存令牌到localStorage
        if (result.data) {
          localStorage.setItem('token', result.data.token);
          localStorage.setItem('token_type', result.data.token_type);
          localStorage.setItem('expires_in', result.data.expires_in.toString());
          localStorage.setItem('login_time', Date.now().toString());
        }
        
        messageApi.success(result.message || '登录成功');
        
        // 延迟跳转，确保消息显示
        setTimeout(() => {
          // 登录成功后重定向到首页
          window.location.href = '/home';
        }, 1000);
      } else {
        messageApi.error(result.message || '登录失败');
      }
    } catch (error) {
      messageApi.error('登录失败，请稍后重试');
      console.error('登录错误:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <App>
      {contextHolder}
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
              name="email"
              rules={[
                { required: true, message: '请输入邮箱' },
                { type: 'email', message: '请输入有效的邮箱地址' }
              ]}
            >
              <Input 
                prefix={<UserOutlined />} 
                placeholder="邮箱" 
                size="large"
              />
            </Form.Item>

            <Form.Item
              name="password"
              rules={[
                { required: true, message: '请输入密码' },
                { min: 6, message: '密码长度不能少于6个字符' }
              ]}
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
    </App>
  );
} 