import type { Route } from "./+types/home";
import { Layout, Menu, Button, Typography } from "antd";
import { UserOutlined, DashboardOutlined, LogoutOutlined } from '@ant-design/icons';
import { useNavigate, useLocation } from 'react-router-dom';

const { Header, Content, Sider } = Layout;
const { Title } = Typography;

export function meta({}: Route.MetaArgs) {
  return [
    { title: "首页 - Go Web MVC 应用" },
    { name: "description", content: "Go Web MVC 应用首页" },
  ];
}

export default function Home() {
  const navigate = useNavigate();
  const location = useLocation();
  
  // 获取当前选中的菜单项
  const getSelectedKey = () => {
    const path = location.pathname;
    if (path.includes('/users')) return 'user';
    return 'dashboard';
  };

  // 处理登出操作
  const handleLogout = () => {
    // 清除登录状态
    localStorage.removeItem('token');
    localStorage.removeItem('token_type');
    localStorage.removeItem('expires_in');
    localStorage.removeItem('login_time');
    
    window.location.href = '/'; // 返回登录页
  };

  // 处理菜单点击
  const handleMenuClick = (key: string) => {
    switch (key) {
      case 'user':
        navigate('/users');
        break;
      case 'dashboard':
        navigate('/home');
        break;
      default:
        break;
    }
  };

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Header style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
        <div className="text-white font-bold text-xl">Go Web MVC</div>
        <Button 
          type="text" 
          icon={<LogoutOutlined />}
          onClick={handleLogout}
          style={{ color: 'white' }}
        >
          退出登录
        </Button>
      </Header>
      <Layout>
        <Sider width={200} style={{ background: '#fff' }}>
          <Menu
            mode="inline"
            selectedKeys={[getSelectedKey()]}
            style={{ height: '100%', borderRight: 0 }}
            onClick={({key}) => handleMenuClick(key)}
          >
            <Menu.Item key="dashboard" icon={<DashboardOutlined />}>
              控制面板
            </Menu.Item>
            <Menu.Item key="user" icon={<UserOutlined />}>
              用户管理
            </Menu.Item>
          </Menu>
        </Sider>
        <Content style={{ padding: 24, margin: 0, minHeight: 280, background: '#fff' }}>
          {location.pathname === '/home' && (
            <>
              <Title level={3}>欢迎使用 Go Web MVC 应用</Title>
              <p>您已成功登录系统。</p>
            </>
          )}
        </Content>
      </Layout>
    </Layout>
  );
}
