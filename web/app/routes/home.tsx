import type { Route } from "./+types/home";
import { Layout, Menu, Button, Typography } from "antd";
import { UserOutlined, DashboardOutlined, LogoutOutlined } from '@ant-design/icons';

const { Header, Content, Sider } = Layout;
const { Title } = Typography;

export function meta({}: Route.MetaArgs) {
  return [
    { title: "首页 - Go Web MVC 应用" },
    { name: "description", content: "Go Web MVC 应用首页" },
  ];
}

export default function Home() {
  // 处理登出操作
  const handleLogout = () => {
    // 这里应该清除登录状态
    window.location.href = '/'; // 返回登录页
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
            defaultSelectedKeys={['dashboard']}
            style={{ height: '100%', borderRight: 0 }}
          >
            <Menu.Item key="dashboard" icon={<DashboardOutlined />}>
              控制面板
            </Menu.Item>
            <Menu.Item key="account" icon={<UserOutlined />}>
              账户管理
            </Menu.Item>
          </Menu>
        </Sider>
        <Content style={{ padding: 24, margin: 0, minHeight: 280, background: '#fff' }}>
          <Title level={3}>欢迎使用 Go Web MVC 应用</Title>
          <p>您已成功登录系统。</p>
        </Content>
      </Layout>
    </Layout>
  );
}
