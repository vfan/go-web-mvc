# Go Web MVC 项目

本项目是一个基于Gin框架的Go语言Web应用，采用多层架构设计模式，展示了分层架构的实现和最佳实践。

## 架构概述

本项目在传统MVC（Model-View-Controller）基础上进行了扩展，采用了更细化的多层架构：

- **模型(Model)**: 进一步拆分为DAO层和Service层，分别处理数据访问和业务逻辑
- **视图(View)**: 在前后端分离的架构中，后端不直接处理视图，而是提供API接口
- **控制器(Controller)**: 处理HTTP请求和响应，调用相应的Service

这种多层架构更好地实现了关注点分离和单一职责原则，提高了代码的可维护性和可扩展性。

## 多层架构详解

### 1. 数据访问层 (DAO - Data Access Object)

位置：`server/dao/`

职责：
- 提供对数据库的CRUD操作
- 封装数据访问细节
- 使用GORM框架与数据库交互
- 只关注数据的存取，不包含业务逻辑

代码示例：
```go
// UserDAO 用户数据访问对象
type UserDAO struct {
    DB *gorm.DB
}

// GetByID 根据ID获取用户
func (dao *UserDAO) GetByID(id int64) (*model.User, error) {
    var user model.User
    err := dao.DB.First(&user, id).Error
    return &user, err
}
```

### 2. 业务逻辑层 (Service)

位置：`server/service/`

职责：
- 实现应用的业务逻辑和规则
- 协调多个DAO操作，实现事务
- 不直接与数据库交互，通过DAO层访问数据
- 处理业务异常和业务规则验证

代码示例：
```go
// UserService 用户服务
type UserService struct {
    userDAO *dao.UserDAO
}

// Login 用户登录
func (s *UserService) Login(email, password string) (*model.User, error) {
    // 验证登录信息
    user, err := s.userDAO.ValidateLogin(email, password)
    if err != nil {
        return nil, err
    }
    
    // 更新最后登录时间
    now := time.Now()
    user.LastLoginTime = &now
    s.userDAO.UpdateLastLoginTime(user.ID)
    
    return user, nil
}
```

### 3. 控制器层 (Controller)

位置：`server/controllers/`

职责：
- 处理HTTP请求和响应
- 解析和验证请求参数
- 调用相应的Service进行处理
- 构造和返回HTTP响应
- 不包含业务逻辑

代码示例：
```go
// UserController 用户控制器
type UserController struct {
    userService *service.UserService
}

// Login 用户登录
func (a *AuthController) Login(c *gin.Context) {
    var req dto.LoginRequest
    
    // 请求参数绑定
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ParamError(c, "无效的请求参数")
        return
    }
    
    // 验证用户是否存在且密码正确
    user, err := a.userService.Login(req.Email, req.Password)
    if err != nil {
        utils.Unauthorized(c, "用户名或密码错误")
        return
    }
    
    // 生成JWT令牌
    token, err := utils.GenerateToken(user)
    if err != nil {
        utils.InternalError(c, "生成令牌失败")
        return
    }
    
    // 返回响应
    utils.SuccessWithMsg(c, "登录成功", dto.LoginResponse{
        Token:     token,
        TokenType: "Bearer",
        ExpiresIn: expiresIn,
    })
}
```

### 4. 数据传输对象 (DTO - Data Transfer Object)

位置：`server/models/dto/`

职责：
- 定义请求和响应的数据结构
- 实现层与层之间的数据传输
- 隔离内部模型和外部接口
- 与前端交互的数据模型

代码示例：
```go
// 登录请求
type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

// 登录响应
type LoginResponse struct {
    Token     string `json:"token"`
    TokenType string `json:"token_type"`
    ExpiresIn int    `json:"expires_in"` // 过期时间，单位：秒
}
```

## 数据流向

一个典型请求的数据流向：

1. 客户端发送HTTP请求到Controller
2. Controller解析请求参数，调用相应的Service方法
3. Service实现业务逻辑，调用DAO进行数据操作
4. DAO与数据库交互，执行CRUD操作
5. 数据按相反顺序返回：DAO → Service → Controller → 客户端

## 依赖注入

本项目采用手动依赖注入的方式管理组件依赖关系：

```go
// 初始化依赖
func initDependencies(db *gorm.DB) *AppDependencies {
    // 初始化DAO
    userDAO := dao.NewUserDAO(db)
    
    // 初始化服务
    userService := service.NewUserService(userDAO)
    
    // 返回依赖集合
    return &AppDependencies{
        DB:          db,
        UserDAO:     userDAO,
        UserService: userService,
    }
}
```

## 多层架构的优势

1. **关注点分离**：每一层只关注自己的职责，降低耦合度
2. **代码复用**：DAO和Service层可以被多个控制器共享
3. **可测试性**：各层可以独立测试，便于单元测试和集成测试
4. **可维护性**：修改一层代码不会影响其他层，提高代码质量
5. **可扩展性**：可以轻松扩展或替换任何一层的实现，满足不断变化的需求

## 配置管理

项目使用 `.env` 和 `.env.local` 文件来管理配置：

- `.env`: 基础配置文件，包含默认配置
- `.env.local`: 本地覆盖配置文件，可以覆盖基础配置中的设置

## 配置加载顺序

1. 首先加载 `.env` 文件中的所有配置
2. 然后加载 `.env.local` 文件中的配置（如果存在），覆盖相同的配置项
3. 如果某个配置在两个文件中都不存在，则使用代码中定义的默认值


## 使用方法

1. 复制 `.env.example` 文件（如果存在）到 `.env` 文件，并设置基础配置
2. 创建 `.env.local` 文件来覆盖本地开发环境需要的配置
3. 确保 `.env.local` 文件已添加到 `.gitignore` 中，避免将本地配置提交到版本控制

> 注意：不要在 `.env` 文件中存储敏感信息（如生产环境密码和密钥），因为它会被提交到版本控制。敏感信息应该放在 `.env.local` 文件中或者使用环境变量注入。 


# JWT 配置

## JWT配置详细说明

### 1. JWT_SECRET_KEY（JWT密钥）

**作用**：这是用于签名和验证JWT令牌的密钥，是JWT安全机制的核心。
- 签名过程：系统使用此密钥对JWT令牌的头部和载荷部分进行加密签名
- 验证过程：系统使用同样的密钥验证令牌签名是否有效

**重要性**：
- 如果密钥泄露，攻击者可以伪造有效的JWT令牌，从而以任何用户身份访问系统
- 密钥强度直接关系到令牌的安全性

**最佳实践**：
- 在生产环境中使用长度至少32字符的随机字符串
- 定期更换密钥
- 永远不要硬编码在源代码中或提交到版本控制系统

### 2. JWT_TOKEN_EXPIRY（令牌过期时间）

**作用**：定义访问令牌的有效期限，单位为小时。
- 我们系统默认设置为24小时
- 过了这个时间，令牌将自动失效，用户需要重新登录

**重要性**：
- 过期机制限制了令牌被盗用的风险窗口期
- 即使令牌被盗，也只在有限时间内有效

**权衡考虑**：
- 时间太短：用户需要频繁登录，影响体验
- 时间太长：安全风险增加，被盗令牌的可用时间延长

### 3. JWT_REFRESH_EXPIRY（刷新令牌过期时间）

**作用**：定义刷新令牌的有效期限，单位为小时。
- 我们系统默认设置为168小时（7天）
- 刷新令牌用于在访问令牌过期后获取新的访问令牌，无需用户重新登录

**实现机制**：
- 当访问令牌过期时，前端可以使用刷新令牌请求新的访问令牌
- 这允许用户保持登录状态更长时间，同时限制单个访问令牌的有效期

**注意**：
- 当前我们系统只实现了访问令牌，没有实现刷新令牌机制
- 但配置已预留，可以在后续开发中实现完整的刷新机制

### 4. JWT_ISSUER（发行者）

**作用**：标识令牌的发行者，通常是应用程序或组织的名称。
- 在令牌的payload部分包含`iss`字段
- 系统验证令牌时会检查此字段，确保令牌来自预期的发行者

**应用场景**：
- 当存在多个系统或服务时，可以区分不同系统发放的令牌
- 有助于防止在多系统环境中令牌被跨系统使用

**安全价值**：
- 增加了一层验证，确保令牌是由可信来源颁发的
- 在微服务架构中特别有用，可以区分不同服务发放的令牌

# 数据库变更管理

## 数据库结构变更处理

当数据库表结构发生变化时（如添加字段、修改字段类型、新增表等），需要同步更新代码中的模型定义。本项目使用GORM Gen工具自动从数据库生成模型代码，简化了这一过程。

## 自动生成模型代码

### 工具说明

项目中的`server/cmd/gen.go`文件是一个用于自动生成数据模型的工具：

```go
// 从数据库表生成所有模型
// g.GenerateAllTable()

// 或者只生成指定表的模型
// 第一个参数是表名，第二个参数是模型名称
g.GenerateModel("users")
g.GenerateModel("universities")
g.GenerateModel("students")
```

### 使用步骤

当数据库结构变更后，按照以下步骤更新代码：

1. 首先确保已经在数据库中完成了表结构的修改
2. 运行gen.go工具生成新的模型代码：

```bash
cd server
go run cmd/gen.go
```

3. 生成的模型代码位于`server/dao/model/`目录下
4. 检查生成的代码，确认字段定义是否符合预期
5. 如果有新增表，需要在`gen.go`中添加对应的`g.GenerateModel("新表名")`
6. 视情况更新相关的DAO、Service和Controller代码，以支持新的字段或表

### 注意事项

1. **不要手动修改生成的模型代码**：这些文件会在下次运行gen.go时被覆盖
2. **业务逻辑放在Service层**：避免在生成的模型上直接添加业务方法
3. **表关联处理**：GORM Gen会自动识别外键关系，生成关联字段
4. **字段类型映射**：检查生成的Go类型是否符合预期，特别是对于枚举类型
5. **生成配置调整**：如果需要修改生成行为（如添加注释、修改类型映射），可以在`gen.go`中的配置部分进行调整

### 完整的数据库变更流程

1. 设计数据库变更（添加字段、修改字段、新增表等）
2. 编写SQL变更脚本并执行
3. 运行gen.go更新模型代码
4. 更新DAO层以支持新的字段或表
5. 更新或创建Service层的业务逻辑
6. 更新或创建Controller层的接口
7. 更新或创建DTO定义
8. 测试新功能
9. 提交代码

 

