# gin-plus

该项目是在Gin框架的基础上添加一些工作中的常用组件，使其更适合项目开发。

## 约定

```shell
.
├── README.md               # 文档说明
├── app                     # 业务代码目录
│   ├── controller    # 控制器层 - 参数校验，数据返回
│   ├── dao           # 数据访问层(Database Access Object)，所有与数据操作相关的操作都会在dao层进行
│   ├── middleware    # HTTP中间件 
│   ├── model         # 模型层，用于存放model对象
│   ├── proto         # proto文件
│   └── service       # 项目核心业务逻辑
├── config                  # 配置文件
├── docs                    # 文档集合(例如Swagger)              
├── go.mod       
├── go.sum
├── main.go                 # 项目入口
├── pkg                     # 项目相关的模块包
├── routes                  # 路由模块(所有的路由都定义在该目录下)
│   └── routes.go
└── scripts
```

## 框架

- Gin: https://github.com/gin-gonic/gin

## 组件

- 配置文件管理组件：viper https://github.com/spf13/viper