import { useEffect, useState } from 'react';
import { Table, Button, Space, Modal, message, Popconfirm } from 'antd';
import { PlusOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons';
import type { ColumnsType } from 'antd/es/table';
import { getUserList, deleteUser } from '../../services/userService';
import type { User } from '../+types/home';
import UserForm from './components/UserForm';
import { useNavigate } from 'react-router-dom';

// 用户管理页面元数据
export function meta() {
  return [
    { title: "用户管理 - Go Web MVC应用" },
    { name: "description", content: "用户管理页面" },
  ];
}

export default function UserManagement() {
  // 状态定义
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [total, setTotal] = useState<number>(0);
  const [current, setCurrent] = useState<number>(1);
  const [pageSize, setPageSize] = useState<number>(10);
  const [modalVisible, setModalVisible] = useState<boolean>(false);
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  const [messageApi, contextHolder] = message.useMessage();
  const navigate = useNavigate();

  // 加载用户数据
  const loadUsers = async () => {
    setLoading(true);
    try {
      const response = await getUserList({ page: current, pageSize });

      
      console.log("response", response)
      if (response.data) {
        setUsers(response.data.list);
        setTotal(response.data.total);
      }
    } catch (error) {
      messageApi.error('获取用户列表失败');
      console.error('获取用户列表错误:', error);
    } finally {
      setLoading(false);
    }
  };

  // 首次加载和分页变化时重新加载数据
  useEffect(() => {
    loadUsers();
  }, [current, pageSize]);

  // 处理删除用户
  const handleDelete = async (id: number) => {
    try {
      await deleteUser(id);
      messageApi.success('删除用户成功');
      loadUsers(); // 重新加载列表
    } catch (error) {
      messageApi.error('删除用户失败');
      console.error('删除用户错误:', error);
    }
  };

  // 处理编辑用户
  const handleEdit = (user: User) => {
    setCurrentUser(user);
    setModalVisible(true);
  };

  // 处理查看用户详情
  const handleView = (id: number) => {
    navigate(`/home/users/${id}`);
  };

  // 处理添加用户
  const handleAdd = () => {
    setCurrentUser(null);
    setModalVisible(true);
  };

  // 表单提交成功后关闭模态框并刷新列表
  const handleFormSuccess = () => {
    setModalVisible(false);
    loadUsers();
  };

  // 表格列定义
  const columns: ColumnsType<User> = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
      width: 80,
    },
    {
      title: '邮箱',
      dataIndex: 'email',
      key: 'email',
      render: (text, record) => (
        <a onClick={() => handleView(record.id)}>{text}</a>
      ),
    },
    {
      title: '角色',
      dataIndex: 'role',
      key: 'role',
      render: (role) => role === 1 ? '管理员' : '普通用户',
    },
    {
      title: '状态',
      dataIndex: 'status',
      key: 'status',
      render: (status) => status === 1 ? '正常' : '禁用',
    },
    {
      title: '最后登录时间',
      dataIndex: 'last_login_time',
      key: 'last_login_time',
      render: (text) => text ? new Date(text).toLocaleString() : '从未登录',
    },
    {
      title: '创建时间',
      dataIndex: 'created_at',
      key: 'created_at',
      render: (text) => new Date(text).toLocaleString(),
    },
    {
      title: '更新时间',
      dataIndex: 'updated_at',
      key: 'updated_at',
      render: (text) => new Date(text).toLocaleString(),
    },
    {
      title: '操作',
      key: 'action',
      width: 180,
      render: (_, record) => (
        <Space size="middle">
          <Button 
            type="text" 
            icon={<EditOutlined />} 
            onClick={() => handleEdit(record)}
          >
            编辑
          </Button>
          <Popconfirm
            title="确定要删除此用户吗?"
            onConfirm={() => handleDelete(record.id)}
            okText="确定"
            cancelText="取消"
          >
            <Button 
              type="text" 
              danger 
              icon={<DeleteOutlined />}
            >
              删除
            </Button>
          </Popconfirm>
        </Space>
      ),
    },
  ];

  return (
    <div>
      {contextHolder}
      <div className="flex justify-between mb-4">
        <h1 className="text-2xl font-bold">用户管理</h1>
        <Button
          type="primary"
          icon={<PlusOutlined />}
          onClick={handleAdd}
        >
          添加用户
        </Button>
      </div>

      <Table
        columns={columns}
        dataSource={users}
        rowKey="id"
        loading={loading}
        pagination={{
          current,
          pageSize,
          total,
          onChange: (page, size) => {
            setCurrent(page);
            if (size) setPageSize(size);
          },
          showSizeChanger: true,
          showTotal: (total) => `共 ${total} 条记录`,
        }}
      />

      <Modal
        title={currentUser ? '编辑用户' : '添加用户'}
        open={modalVisible}
        onCancel={() => setModalVisible(false)}
        footer={null}
        destroyOnHidden
        
      >
        <UserForm
          user={currentUser}
          onSuccess={handleFormSuccess}
          onCancel={() => setModalVisible(false)}
        />
      </Modal>
    </div>
  );
} 