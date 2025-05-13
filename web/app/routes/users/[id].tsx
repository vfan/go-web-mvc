import { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Card, Button, Descriptions, Spin, message } from 'antd';
import { ArrowLeftOutlined } from '@ant-design/icons';
import { getUserDetail } from '../../services/userService';
import type { User } from '../+types/home';

// 用户详情页面元数据
export function meta({ params }: { params: { id: string } }) {
  return [
    { title: `用户详情 - 用户ID: ${params.id} - Go Web MVC应用` },
    { name: "description", content: "用户详情页面" },
  ];
}

export default function UserDetail() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [messageApi, contextHolder] = message.useMessage();

  useEffect(() => {
    const fetchUserDetail = async () => {
      if (!id) return;
      
      setLoading(true);
      try {
        const response = await getUserDetail(parseInt(id));
        if (response.data) {
          setUser(response.data);
        } else {
          messageApi.error('获取用户详情失败');
        }
      } catch (error) {
        messageApi.error('获取用户详情失败');
        console.error('获取用户详情错误:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchUserDetail();
  }, [id, messageApi]);

  // 返回用户列表页
  const handleBack = () => {
    navigate('/users');
  };

  return (
    <div className="p-6">
      {contextHolder}
      <Button 
        type="text" 
        icon={<ArrowLeftOutlined />} 
        onClick={handleBack}
        className="mb-4"
      >
        返回用户列表
      </Button>

      <Card title="用户详情" className="shadow-sm">
        {loading ? (
          <div className="flex justify-center items-center p-12">
            <Spin size="large" />
          </div>
        ) : user ? (
          <Descriptions bordered column={1}>
            <Descriptions.Item label="ID">{user.id}</Descriptions.Item>
            <Descriptions.Item label="用户名">{user.username}</Descriptions.Item>
            <Descriptions.Item label="邮箱">{user.email}</Descriptions.Item>
            <Descriptions.Item label="角色">{user.role}</Descriptions.Item>
            <Descriptions.Item label="创建时间">
              {new Date(user.created_at).toLocaleString()}
            </Descriptions.Item>
            <Descriptions.Item label="更新时间">
              {new Date(user.updated_at).toLocaleString()}
            </Descriptions.Item>
          </Descriptions>
        ) : (
          <div className="text-center p-8 text-gray-500">
            未找到用户信息
          </div>
        )}
      </Card>
    </div>
  );
} 