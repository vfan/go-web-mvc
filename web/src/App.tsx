import { useState, useEffect } from 'react'
import { Layout, Menu, Avatar, Dropdown, theme } from 'antd'
import { 
  UserOutlined, 
  MenuFoldOutlined, 
  MenuUnfoldOutlined, 
  HomeOutlined,
  TeamOutlined,
  SettingOutlined,  
  LogoutOutlined,
  BankOutlined
} from '@ant-design/icons'
import { Outlet, Link, useNavigate, useLocation } from 'react-router-dom'
import './App.css'
import api from './utils/api'

const { Header, Sider, Content } = Layout

function App() {
  const [collapsed, setCollapsed] = useState(false)
  const [userEmail, setUserEmail] = useState('')
  const navigate = useNavigate()
  const location = useLocation()
  const {
    token: { colorBgContainer },
  } = theme.useToken()

  // 获取用户信息
  useEffect(() => {
    const userInfoStr = localStorage.getItem('userInfo');
    if (userInfoStr) {
      try {
        const userInfo = JSON.parse(userInfoStr);
        setUserEmail(userInfo.email || '');
      } catch (error) {
        console.error('解析用户信息失败:', error);
      }
    }
  }, []);

  // 定义菜单项
  const items = [
    {
      key: '/',
      path: '/',
      icon: <HomeOutlined />,
      label: <Link to="/">首页</Link>,
    },
    {
      key: '/user',
      path: '/user',
      icon: <TeamOutlined />,
      label: <Link to="/user">用户管理</Link>,
    },
    {
      key: '/university',
      path: '/university',
      icon: <BankOutlined />,
      label: <Link to="/university">大学管理</Link>,
    },
    {
      key: '/settings',
      path: '/settings',
      icon: <SettingOutlined />,
      label: <Link to="/settings">系统设置</Link>,
    },
    {
      key: '/router',
      path: '/router',
      icon: <SettingOutlined />,
      label: <Link to="/router">路由管理</Link>,
    },
  ];

  // 获取当前选中的菜单项的键值
  const selectedKey = items.find(item => 
    location.pathname === item.path || 
    (location.pathname !== '/' && location.pathname.startsWith(item.path))
  )?.key || '/';

  const userMenuItems = [
    {
      key: 'profile',
      icon: <UserOutlined />,
      label: '个人信息',
    },
    {
      key: 'logout',
      icon: <LogoutOutlined />,
      label: '退出登录',
      onClick: () => {
        handleLogout();
      },
    },
  ]

  // 处理登出
  const handleLogout = async () => {
    try {
      // 调用后端登出接口，清除Cookie
      await api.post('/auth/logout');
    } catch (error) {
      console.error('登出请求失败:', error);
    }
    
    // 清除本地存储的用户信息
    localStorage.removeItem('userInfo');
    navigate('/login');
  };

  return (
    <Layout style={{ minHeight: '100vh', width: '100%' }}>
      <Sider 
        trigger={null} 
        collapsible 
        collapsed={collapsed}
        theme="light"
        style={{ boxShadow: '0 2px 8px rgba(0, 0, 0, 0.15)' }}
      >
        <div className="p-4" style={{ textAlign: 'left' }}>
          <h2 className="text-2xl  font-bold">学生管理系统</h2>
        </div>
        <Menu
          theme="light"
          mode="inline"
          selectedKeys={[selectedKey]}
          items={items}
        />
      </Sider>
      <Layout>
        <Header style={{ padding: 0, background: colorBgContainer, boxShadow: '0 1px 4px rgba(0, 0, 0, 0.08)' }}>
          <div className="flex justify-between items-center px-4">
            <div>
              {collapsed ? (
                <MenuUnfoldOutlined
                  className="text-xl cursor-pointer"
                  onClick={() => setCollapsed(!collapsed)}
                />
              ) : (
                <MenuFoldOutlined
                  className="text-xl cursor-pointer"
                  onClick={() => setCollapsed(!collapsed)}
                />
              )}
            </div>
            <div className="flex items-center">
            
              <Dropdown menu={{ items: userMenuItems }} placement="bottomRight">
                <div className="flex items-center cursor-pointer">
                  <Avatar 
                    icon={<UserOutlined />} 
                    style={{ marginRight: '8px' }}
                  />
                  <span>{userEmail || '用户'}</span>
                </div>
              </Dropdown>
            </div>
          </div>
        </Header>
        <Content
          style={{
            margin: '24px 16px',
            padding: 24,
            background: colorBgContainer,
            borderRadius: 8,
            minHeight: 280,
            overflow: 'auto'
          }}
        >
          <div>Hello router</div>
          <Outlet />
        </Content>
      </Layout>
    </Layout>
  )
}

export default App
