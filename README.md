# Virtual Campus Tour 2.0 Backend

基于 Go 语言和 Gin 框架开发的虚拟校园导览系统后端服务。

## 项目结构

```
.
├── config/              # 配置文件
├── docs/                # 文档
├── internal/            # 内部代码
│   ├── handler/         # HTTP 处理器
│   ├── middleware/      # 中间件
│   ├── model/           # 数据模型
│   ├── repository/      # 数据访问层
│   └── service/         # 业务逻辑层
├── pkg/                 # 公共代码包
│   ├── database/        # 数据库相关
│   ├── logger/          # 日志相关
│   └── utils/           # 工具函数
├── main.go              # 应用程序入口
├── go.mod               # Go 模块文件
└── README.md            # 项目说明文档
```

## 技术栈

- Go
- Gin Web 框架
- GORM
- MySQL
- JWT 认证

## 快速开始

1. 克隆项目
2. 安装依赖：`go mod tidy`
3. 配置数据库：修改 `config/config.yaml` 中的数据库配置
4. 运行项目：`go run main.go`

## API 文档

API 文档位于 `docs` 目录下。