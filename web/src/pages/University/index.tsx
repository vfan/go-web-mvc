import { useState, useEffect } from 'react';
import { Table, Button, Space, Typography, Modal, Input, Form, App } from 'antd';
import { PlusOutlined, EditOutlined, DeleteOutlined, ExclamationCircleOutlined, SyncOutlined } from '@ant-design/icons';
import { getUniversityList, createUniversity, updateUniversity, deleteUniversity } from '../../api/university';
import type { University, UniversityCreateParams, UniversityUpdateParams } from '../../api/university';

const { Title } = Typography;
const { Search } = Input;

function UniversityList() {
  const { message, modal } = App.useApp();
  const [universities, setUniversities] = useState<University[]>([]);
  const [loading, setLoading] = useState(false);
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [currentUniversity, setCurrentUniversity] = useState<University | null>(null);
  const [form] = Form.useForm();
  const [searchText, setSearchText] = useState('');
  const [pagination, setPagination] = useState({
    current: 1,
    pageSize: 10,
    total: 0
  });

  // 加载大学数据
  const fetchUniversities = async (page = 1, pageSize = 10) => {
    setLoading(true);
    try {
      const response = await getUniversityList(page, pageSize);
      console.log(response.list);
      setUniversities(response.list || []);
      setPagination({
        ...pagination,
        current: page,
        total: response.total || 0
      });
    } catch (error) {
      console.error('获取大学列表失败:', error);
      setUniversities([]);
    } finally {
      setLoading(false);
    }
  };

  // 初始加载
  useEffect(() => {
    fetchUniversities(pagination.current, pagination.pageSize);
  }, []);

  // 处理表格分页变化
  const handleTableChange = (pagination: { current?: number; pageSize?: number }) => {
    fetchUniversities(pagination.current || 1, pagination.pageSize || 10);
  };

  // 根据搜索文本过滤大学
  const filteredUniversities = (universities || []).filter(university => 
    university.name?.toLowerCase().includes(searchText.toLowerCase())
  );

  // 处理添加大学
  const handleAddUniversity = async (values: UniversityCreateParams) => {
    try {
      await createUniversity(values);
      message.success('大学已添加');
      setIsModalVisible(false);
      form.resetFields();
      fetchUniversities(pagination.current, pagination.pageSize);
    } catch (error) {
      if (error instanceof Error) {
        message.error(`添加大学失败: ${error.message}`);
      } else {
        message.error('添加大学失败，请重试');
      }
    }
  };

  // 处理编辑大学
  const handleEditUniversity = async (id: number, values: UniversityUpdateParams) => {
    try {
      await updateUniversity(id, values);
      message.success('大学已更新');
      setIsModalVisible(false);
      form.resetFields();
      fetchUniversities(pagination.current, pagination.pageSize);
    } catch (error) {
      if (error instanceof Error) {
        message.error(`更新大学失败: ${error.message}`);
      } else {
        message.error('更新大学失败，请重试');
      }
    }
  };

  // 处理添加/编辑表单提交
  const handleAddOrEditUniversity = () => {
    form.validateFields().then((values: UniversityCreateParams) => {
      if (currentUniversity) {
        // 编辑现有大学
        handleEditUniversity(currentUniversity.id, values);
      } else {
        // 添加新大学
        handleAddUniversity(values);
      }
    });
  };

  // 处理删除大学
  const handleDeleteUniversity = (id: number) => {
    modal.confirm({
      title: '确认删除',
      icon: <ExclamationCircleOutlined />,
      content: '您确定要删除此大学吗？此操作不可逆。',
      onOk: async () => {
        try {
          await deleteUniversity(id);
          message.success('大学已删除');
          fetchUniversities(pagination.current, pagination.pageSize);
        } catch (error) {
          if (error instanceof Error) {
            message.error(`删除大学失败: ${error.message}`);
          } else {
            message.error('删除大学失败，请重试');
          }
        }
      },
    });
  };

  // 打开编辑模态框
  const handleEditClick = (university: University) => {
    setCurrentUniversity(university);
    form.setFieldsValue({
      name: university.name
    });
    setIsModalVisible(true);
  };

  // 打开添加模态框
  const handleAddClick = () => {
    setCurrentUniversity(null);
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
    fetchUniversities(pagination.current, pagination.pageSize);
  };

  // 表格列定义
  const columns = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
    },
    {
      title: '大学名称',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: '创建时间',
      dataIndex: 'created_at',
      key: 'created_at',
      render: (time: string) => new Date(time).toLocaleString(),
    },
    {
      title: '更新时间',
      dataIndex: 'updated_at',
      key: 'updated_at',
      render: (time: string) => new Date(time).toLocaleString(),
    },
    {
      title: '操作',
      key: 'action',
      render: (_: unknown, record: University) => (
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
            onClick={() => handleDeleteUniversity(record.id)}
          >
            删除
          </Button>
        </Space>
      ),
    },
  ];

  return (
    <div className="p-4">
      <Title level={2}>大学管理</Title>
      
      <div className="flex justify-between mb-4">
        <div className="flex gap-2">
          <Search
            placeholder="搜索大学名称"
            allowClear
            style={{ width: 300 }}
            onSearch={(value: string) => setSearchText(value)}
            onChange={(e: React.ChangeEvent<HTMLInputElement>) => setSearchText(e.target.value)}
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
          添加大学
        </Button>
      </div>
      
      <Table
        rowKey="id"
        columns={columns}
        dataSource={filteredUniversities}
        pagination={{
          ...pagination,
          showSizeChanger: true,
          showQuickJumper: true,
          showTotal: (total: number) => `共 ${total} 条记录`
        }}
        loading={loading}
        onChange={handleTableChange}
      />
      
      <Modal
        title={currentUniversity ? '编辑大学' : '添加大学'}
        open={isModalVisible}
        onOk={handleAddOrEditUniversity}
        onCancel={handleModalCancel}
        okText={currentUniversity ? '更新' : '添加'}
        cancelText="取消"
      >
        <Form
          form={form}
          layout="vertical"
          name="universityForm"
        >
          <Form.Item
            name="name"
            label="大学名称"
            rules={[{ required: true, message: '请输入大学名称' }]}
          >
            <Input placeholder="请输入大学名称" />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
}

export default UniversityList; 