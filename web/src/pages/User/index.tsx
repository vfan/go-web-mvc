import { useState, useEffect } from 'react';
import { Table, Button, Space, Card, Typography, Modal, Input, Form, Select, App } from 'antd';
import { PlusOutlined, EditOutlined, DeleteOutlined, ExclamationCircleOutlined, SyncOutlined } from '@ant-design/icons';
import { getUserList, createUser, updateUser, deleteUser } from '../../api/user';
import type { User, UserCreateParams, UserUpdateParams } from '../../api/user';

const { Title } = Typography;
const { Search } = Input;
const { Option } = Select;

// 角色映射
const roleMap: Record<number, string> = {
  1: '管理员',
  2: '普通用户'
};

// 状态映射
const statusMap: Record<number, string> = {
  1: '正常',
  0: '禁用'
};

function UserList() {
  const { message, modal } = App.useApp();
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(false);
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  const [form] = Form.useForm();
  const [searchText, setSearchText] = useState('');
  const [pagination, setPagination] = useState({
    current: 1,
    pageSize: 10,
    total: 0
  });

  // 加载用户数据
  const fetchUsers = async (page = 1, pageSize = 10) => {
    setLoading(true);
    try {
      const response = await getUserList(page, pageSize);
      console.log(response);
      setUsers(response.list || []);
      setPagination({
        ...pagination,
        current: page,
        total: response.total || 0
      });
    } catch (error) {
      // 错误已在API服务中处理
      console.error('获取用户列表失败:', error);
      setUsers([]);
    } finally {
      setLoading(false);
    }
  };

  // 初始加载
  useEffect(() => {
    fetchUsers(pagination.current, pagination.pageSize);
  }, []);

  // 处理表格分页变化
  const handleTableChange = (pagination: any) => {
    fetchUsers(pagination.current, pagination.pageSize);
  };

  // 根据搜索文本过滤用户
  const filteredUsers = (users || []).filter(user => 
    user.email?.toLowerCase().includes(searchText.toLowerCase())
  );

  // 处理添加用户
  const handleAddUser = async (values: UserCreateParams) => {
    try {
      await createUser(values);
      message.success('用户已添加');
      setIsModalVisible(false);
      form.resetFields();
      fetchUsers(pagination.current, pagination.pageSize);
    } catch (error) {
      // 在界面上显示错误提示
      if (error instanceof Error) {
        message.error(`添加用户失败: ${error.message}`);
      } else {
        message.error('添加用户失败，请重试');
      }
    }
  };

  // 处理编辑用户
  const handleEditUser = async (id: number, values: UserUpdateParams) => {
    try {
      await updateUser(id, values);
      message.success('用户已更新');
      setIsModalVisible(false);
      form.resetFields();
      fetchUsers(pagination.current, pagination.pageSize);
    } catch (error) {
      // 在界面上显示错误提示
      if (error instanceof Error) {
        message.error(`更新用户失败: ${error.message}`);
      } else {
        message.error('更新用户失败，请重试');
      }
    }
  };

  // 处理添加/编辑表单提交
  const handleAddOrEditUser = () => {
    form.validateFields().then(values => {
      if (currentUser) {
        // 编辑现有用户
        handleEditUser(currentUser.id, values);
      } else {
        // 添加新用户
        handleAddUser(values);
      }
    });
  };

  // 处理删除用户
  const handleDeleteUser = (id: number) => {
    modal.confirm({
      title: '确认删除',
      icon: <ExclamationCircleOutlined />,
      content: '您确定要删除此用户吗？此操作不可逆。',
      onOk: async () => {
        try {
          await deleteUser(id);
          message.success('用户已删除');
          fetchUsers(pagination.current, pagination.pageSize);
        } catch (error) {
          // 在界面上显示错误提示
          if (error instanceof Error) {
            message.error(`删除用户失败: ${error.message}`);
          } else {
            message.error('删除用户失败，请重试');
          }
        }
      },
    });
  };

  // 打开编辑模态框
  const handleEditClick = (user: User) => {
    setCurrentUser(user);
    form.setFieldsValue({
      email: user.email,
      role: user.role,
      status: user.status
    });
    setIsModalVisible(true);
  };

  // 打开添加模态框
  const handleAddClick = () => {
    setCurrentUser(null);
    form.resetFields();
    setIsModalVisible(true);
  };

  // 关闭模态框
  const handleModalCancel = () => {
    setIsModalVisible(false);
    form.resetFields();
  };

  // 刷新数据
  const handleRefresh = () => {
    fetchUsers(pagination.current, pagination.pageSize);
  };

  // 表格列定义
  const columns = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
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
      render: (role: number) => roleMap[role] || '未知角色',
    },
    {
      title: '状态',
      dataIndex: 'status',
      key: 'status',
      render: (status: number) => statusMap[status] || '未知状态',
    },
    {
      title: '最后登录时间',
      dataIndex: 'last_login_time',
      key: 'last_login_time',
      render: (time: string) => time ? new Date(time).toLocaleString() : '从未登录',
    },
    {
      title: '创建时间',
      dataIndex: 'created_at',
      key: 'created_at',
      render: (time: string) => new Date(time).toLocaleString(),
    },
    {
      title: '操作',
      key: 'action',
      render: (_: any, record: User) => (
        <Space size="middle">
          <Button 
            type="primary" 
            icon={<EditOutlined />} 
            onClick={() => handleEditClick(record)}
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
        <div className="flex gap-2">
          <Search
            placeholder="搜索邮箱"
            allowClear
            style={{ width: 300 }}
            onSearch={value => setSearchText(value)}
            onChange={e => setSearchText(e.target.value)}
          />
          <Button 
            icon={<SyncOutlined />} 
            onClick={handleRefresh}
          >
            刷新
          </Button>
        </div>
        <Button 
          type="primary" 
          icon={<PlusOutlined />} 
          onClick={handleAddClick}
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
          pagination={{
            ...pagination,
            showSizeChanger: true,
            showTotal: (total) => `共 ${total} 条记录`
          }}
          onChange={handleTableChange}
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
            name="email"
            label="邮箱"
            rules={[
              { required: true, message: '请输入邮箱' },
              { type: 'email', message: '请输入有效的邮箱地址' }
            ]}
          >
            <Input />
          </Form.Item>
          {!currentUser && (
            <Form.Item
              name="password"
              label="密码"
              rules={[{ required: true, message: '请输入密码' }]}
            >
              <Input.Password />
            </Form.Item>
          )}
          <Form.Item
            name="role"
            label="角色"
            rules={[{ required: true, message: '请选择角色' }]}
          >
            <Select placeholder="请选择角色">
              <Option value={1}>管理员</Option>
              <Option value={2}>普通用户</Option>
            </Select>
          </Form.Item>
          <Form.Item
            name="status"
            label="状态"
            rules={[{ required: true, message: '请选择状态' }]}
          >
            <Select placeholder="请选择状态">
              <Option value={1}>正常</Option>
              <Option value={0}>禁用</Option>
            </Select>
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
}

export default UserList; 