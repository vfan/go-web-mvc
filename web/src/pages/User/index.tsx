import { useState, useEffect } from 'react';
import { Table, Button, Space, Card, Typography, Modal, message, Input, Form } from 'antd';
import { PlusOutlined, EditOutlined, DeleteOutlined, ExclamationCircleOutlined } from '@ant-design/icons';

const { Title } = Typography;
const { confirm } = Modal;
const { Search } = Input;

interface User {
  id: number;
  username: string;
  email: string;
  role: string;
  createdAt: string;
}

// 模拟用户数据
const mockUsers: User[] = [
  { id: 1, username: 'admin', email: 'admin@example.com', role: '管理员', createdAt: '2023-01-01' },
  { id: 2, username: 'user1', email: 'user1@example.com', role: '普通用户', createdAt: '2023-01-02' },
  { id: 3, username: 'user2', email: 'user2@example.com', role: '普通用户', createdAt: '2023-01-03' },
];

function UserList() {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(false);
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  const [form] = Form.useForm();
  const [searchText, setSearchText] = useState('');

  // 加载用户数据
  useEffect(() => {
    fetchUsers();
  }, []);

  // 模拟获取用户数据
  const fetchUsers = () => {
    setLoading(true);
    // 在实际项目中应该从API获取数据
    setTimeout(() => {
      setUsers(mockUsers);
      setLoading(false);
    }, 500);
  };

  // 根据搜索文本过滤用户
  const filteredUsers = users.filter(user => 
    user.username.toLowerCase().includes(searchText.toLowerCase()) || 
    user.email.toLowerCase().includes(searchText.toLowerCase())
  );

  // 处理添加/编辑用户
  const handleAddOrEditUser = () => {
    form.validateFields().then(values => {
      if (currentUser) {
        // 编辑现有用户
        const updatedUsers = users.map(user => 
          user.id === currentUser.id ? { ...user, ...values } : user
        );
        setUsers(updatedUsers);
        message.success('用户已更新');
      } else {
        // 添加新用户
        const newUser = {
          id: Math.max(...users.map(u => u.id), 0) + 1,
          ...values,
          createdAt: new Date().toISOString().split('T')[0]
        };
        setUsers([...users, newUser]);
        message.success('用户已添加');
      }
      handleModalCancel();
    });
  };

  // 处理删除用户
  const handleDeleteUser = (id: number) => {
    confirm({
      title: '确认删除',
      icon: <ExclamationCircleOutlined />,
      content: '您确定要删除此用户吗？此操作不可逆。',
      onOk() {
        const updatedUsers = users.filter(user => user.id !== id);
        setUsers(updatedUsers);
        message.success('用户已删除');
      },
    });
  };

  // 打开编辑模态框
  const handleEditUser = (user: User) => {
    setCurrentUser(user);
    form.setFieldsValue(user);
    setIsModalVisible(true);
  };

  // 打开添加模态框
  const handleAddUser = () => {
    setCurrentUser(null);
    form.resetFields();
    setIsModalVisible(true);
  };

  // 关闭模态框
  const handleModalCancel = () => {
    setIsModalVisible(false);
    form.resetFields();
  };

  // 表格列定义
  const columns = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
    },
    {
      title: '用户名',
      dataIndex: 'username',
      key: 'username',
    },
    {
      title: '邮箱',
      dataIndex: 'email',
      key: 'email',
    },
    {
      title: '角色',
      dataIndex: 'role',
      key: 'role',
    },
    {
      title: '创建时间',
      dataIndex: 'createdAt',
      key: 'createdAt',
    },
    {
      title: '操作',
      key: 'action',
      render: (_: any, record: User) => (
        <Space size="middle">
          <Button 
            type="primary" 
            icon={<EditOutlined />} 
            onClick={() => handleEditUser(record)}
          >
            编辑
          </Button>
          <Button 
            danger 
            icon={<DeleteOutlined />} 
            onClick={() => handleDeleteUser(record.id)}
          >
            删除
          </Button>
        </Space>
      ),
    },
  ];

  return (
    <div className="p-4">
      <Title level={2}>用户管理</Title>
      
      <div className="flex justify-between mb-4">
        <Search
          placeholder="搜索用户名或邮箱"
          allowClear
          style={{ width: 300 }}
          onSearch={value => setSearchText(value)}
          onChange={e => setSearchText(e.target.value)}
        />
        <Button 
          type="primary" 
          icon={<PlusOutlined />} 
          onClick={handleAddUser}
        >
          添加用户
        </Button>
      </div>
      
      <Card>
        <Table 
          columns={columns} 
          dataSource={filteredUsers} 
          rowKey="id" 
          loading={loading}
          pagination={{ pageSize: 10 }}
        />
      </Card>

      {/* 添加/编辑用户模态框 */}
      <Modal
        title={currentUser ? '编辑用户' : '添加用户'}
        open={isModalVisible}
        onOk={handleAddOrEditUser}
        onCancel={handleModalCancel}
      >
        <Form form={form} layout="vertical">
          <Form.Item
            name="username"
            label="用户名"
            rules={[{ required: true, message: '请输入用户名' }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            name="email"
            label="邮箱"
            rules={[
              { required: true, message: '请输入邮箱' },
              { type: 'email', message: '请输入有效的邮箱地址' }
            ]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            name="role"
            label="角色"
            rules={[{ required: true, message: '请选择角色' }]}
          >
            <Input />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
}

export default UserList; 