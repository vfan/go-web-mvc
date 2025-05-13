import { Card, Typography, Statistic, Row, Col } from 'antd';
import { UserOutlined, FileOutlined, ClockCircleOutlined } from '@ant-design/icons';
import React from 'react';

const { Title } = Typography;

// 定义Home组件
const Home: React.FC = () => {
  return (
    <div className="p-4">
      <Title level={2}>系统概览</Title>
      
      <Row gutter={16} className="mt-4">
        <Col span={8}>
          <Card>
            <Statistic
              title="总用户数"
              value={128}
              prefix={<UserOutlined />}
            />
          </Card>
        </Col>
        <Col span={8}>
          <Card>
            <Statistic
              title="系统数据量"
              value={256}
              prefix={<FileOutlined />}
              suffix="条"
            />
          </Card>
        </Col>
        <Col span={8}>
          <Card>
            <Statistic
              title="运行时间"
              value={30}
              prefix={<ClockCircleOutlined />}
              suffix="天"
            />
          </Card>
        </Col>
      </Row>
      
      <Card className="mt-4">
        <p>欢迎使用 Go Web MVC 后台管理系统</p>
        <p>今天是 {new Date().toLocaleDateString()}</p>
      </Card>
    </div>
  );
};

export default Home; 