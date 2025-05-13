import { Card, Typography } from 'antd';

const { Title } = Typography;

function Home() {
  return (
    <div className="p-4">
      <Title level={2}>首页</Title>
      <Card className="mt-4">
        <p>欢迎使用 Go Web MVC 系统</p>
      </Card>
    </div>
  );
}

export default Home; 