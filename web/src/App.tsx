import { useState } from 'react'
import { Layout, Menu, Avatar, Dropdown, theme } from 'antd'
import { UserOutlined, MenuFoldOutlined, MenuUnfoldOutlined } from '@ant-design/icons'
import { Outlet, Link, useNavigate } from 'react-router-dom'
import './App.css'

const { Header, Sider, Content } = Layout

function App() {
  const [collapsed, setCollapsed] = useState(false)
  const navigate = useNavigate()
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken()

  const items = [
    {
      key: '1',
      label: <Link to="/">首页</Link>,
    },
    // 可以添加其他菜单项
  ]

  const userMenuItems = [
    {
      key: 'profile',
      label: '个人信息',
    },
    {
      key: 'logout',
      label: '退出登录',
      onClick: () => {
        localStorage.removeItem('token')
        navigate('/login')
      },
    },
  ]

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider 
        trigger={null} 
        collapsible 
        collapsed={collapsed}
        theme="light"
      >
        <div className="text-center p-4">
          <h1 className="text-xl font-bold">Go Web MVC</h1>
        </div>
        <Menu
          theme="light"
          mode="inline"
          defaultSelectedKeys={['1']}
          items={items}
        />
      </Sider>
      <Layout>
        <Header style={{ padding: 0, background: colorBgContainer }}>
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
            <div>
              <Dropdown menu={{ items: userMenuItems }} placement="bottomRight">
                <Avatar 
                  icon={<UserOutlined />} 
                  className="cursor-pointer"
                />
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
          }}
        >
          <Outlet />
        </Content>
      </Layout>
    </Layout>
  )
}

export default App
