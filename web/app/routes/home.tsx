import type { Route } from "./+types/home";
import { Welcome } from "../welcome/welcome";
import { Button, Space, Typography, Card, Divider } from "antd";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Go Web MVC 应用" },
    { name: "description", content: "基于 React 和 Ant Design 构建的应用" },
  ];
}

export default function Home() {
  return (
    <div className="p-8">
      <Typography.Title level={2}>欢迎使用 Ant Design 组件</Typography.Title>
      <Divider />
      <Card title="组件示例" style={{ width: 500 }}>
        <Space direction="vertical">
          <Button type="primary">主要按钮</Button>
          <Button>默认按钮</Button>
          <Button type="dashed">虚线按钮</Button>
          <Button type="link">链接按钮</Button>
        </Space>
      </Card>
      
      <Divider />
      <Welcome />
    </div>
  );
}
