import { useState } from 'react'
import { Layout, Menu, Avatar, Dropdown, theme, Badge, Tooltip } from 'antd'
import { 
  UserOutlined, 
  MenuFoldOutlined, 
  MenuUnfoldOutlined, 
  HomeOutlined,
  TeamOutlined,
  SettingOutlined,
  BellOutlined,
  LogoutOutlined
} from '@ant-design/icons'
import { Outlet, Link, useNavigate, useLocation } from 'react-router-dom'
import './App.css'

const { Header, Sider, Content } = Layout

function App() {
  const [collapsed, setCollapsed] = useState(false)
  const navigate = useNavigate()
  const location = useLocation()
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken()

  // 获取当前选中的菜单项
  const getSelectedKey = () => {
    if (location.pathname === '/') return '1'
    if (location.pathname.startsWith('/user')) return '2'
    if (location.pathname.startsWith('/settings')) return '3'
    return '1'
  }

  const items = [
    {
      key: '1',
      icon: <HomeOutlined />,
      label: <Link to="/">首页</Link>,
    },
    {
      key: '2',
      icon: <TeamOutlined />,
      label: <Link to="/user">用户管理</Link>,
    },
    {
      key: '3',
      icon: <SettingOutlined />,
      label: <Link to="/settings">系统设置</Link>,
    },
  ]

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
        localStorage.removeItem('token')
        navigate('/login')
      },
    },
  ]

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
          <h1 className="text-xl font-bold">Go Web MVC</h1>
        </div>
        <Menu
          theme="light"
          mode="inline"
          selectedKeys={[getSelectedKey()]}
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
              <Tooltip title="消息通知">
                <Badge count={5} size="small">
                  <BellOutlined style={{ fontSize: '20px', marginRight: '24px', cursor: 'pointer' }} />
                </Badge>
              </Tooltip>
              <Dropdown menu={{ items: userMenuItems }} placement="bottomRight">
                <div className="flex items-center cursor-pointer">
                  <Avatar 
                    icon={<UserOutlined />} 
                    style={{ marginRight: '8px' }}
                  />
                  <span>管理员</span>
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
            borderRadius: borderRadiusLG,
            minHeight: 280,
            overflow: 'auto'
          }}
        >
          <Outlet />
        </Content>
      </Layout>
    </Layout>
  )
}

export default App
