# 数据访问层（DAO）

## 目录结构

```
server/dao/
├── model/        # 自动生成的数据模型
│   └── *.gen.go  # GORM 生成的模型文件
├── *_dao.go      # 各实体的DAO实现
└── README.md     # 本文档
```

## 设计说明

1. `model` 目录存放GORM自动生成的数据模型，这些模型与数据库表结构一一对应
2. 各个实体的DAO文件（如 `user_dao.go`）提供针对特定数据模型的CRUD操作
3. DAO层只负责数据访问，不包含业务逻辑

## 使用方法

DAO层通过依赖注入的方式被Service层使用：

```go
// 创建DAO实例
userDAO := dao.NewUserDAO(db)

// 创建服务并注入DAO
userService := service.NewUserService(userDAO)
```

## 与原有models的区别

1. 自动生成模型：使用 gorm.io/gen 工具自动生成与数据库表结构一致的模型
2. 分层清晰：模型定义和数据访问逻辑分离
3. 易于扩展：可以针对不同数据源创建不同的DAO实现
4. 代码复用：通过接口抽象，方便测试和模拟 