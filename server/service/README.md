# 服务层（Service）

## 目录结构

```
server/service/
├── *_service.go  # 各实体的服务实现
└── README.md     # 本文档
```

## 设计说明

1. 服务层位于Controller和DAO之间，处理业务逻辑
2. 服务层通过依赖注入方式使用DAO层，解耦数据访问和业务逻辑
3. 服务层实现了业务规则和数据转换，提供面向业务的API给控制器使用

## 使用方法

Service层通过依赖注入的方式被Controller层使用：

```go
// 创建服务实例
userService := service.NewUserService(userDAO)

// 创建控制器并注入服务
userController := controllers.NewUserController(userService)
```

## 作用和职责

1. 实现业务规则和流程
2. 协调多个DAO操作，实现事务
3. 数据校验和处理
4. 处理业务异常
5. 将数据模型转换为业务模型或DTO

## 与原有模型层的区别

1. 关注点分离：Service层只关注业务逻辑，不直接与数据库交互
2. 可测试性：便于单元测试和模拟
3. 可复用性：业务逻辑可在不同控制器中复用
4. 维护性：业务规则集中在一处，便于维护和更新 