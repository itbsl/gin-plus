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
│   ├── config.dev.yaml
│   ├── config.go
│   ├── config.release.yaml
│   └── config.test.yaml
├── docs                    # 文档集合(例如Swagger) 
├── global
│   └── global.go     # 全局变量             
├── go.mod       
├── go.sum
├── main.go                 # 项目入口
├── pkg                     # 项目相关的模块包
│   └── setting
│       ├── setting.go       # 配置文件读取
│       ├── setting_test.go  # 单元测试文件
│       └── setting_test.yaml # 单元测试用例
├── routes                  # 路由模块(所有的路由都定义在该目录下)
│   └── routes.go
└── scripts                 # 脚本相关文件
```

## 使用
```shell
git clone https://github.com/itbsl/gin-plus.git
cd gin-plus
go mod tidy
go build
./gin-plus -f ./config/config.dev.yaml
```
## 框架

- Gin: https://github.com/gin-gonic/gin

## 组件

- 配置文件管理组件：https://github.com/spf13/viper <br>
- 日志库： https://github.com/uber-go/zap
- 日志切割：https://github.com/natefinch/lumberjack


### viper

```shell
读取配置文件的代码在 gin-plus/pkg/settting/setting.go文件
1.优先通过命令行参数 -f 获取显式指定的配置文件路径读取，例如 ./gin-plus -f ./etc/config.yaml
2.如果在命令指定 -f 参数，则尝试通过环境变量 gin_plus_mode 获取是 dev还是release，根据此值来决定读取那个配置文件
3.如果命令行参数和环境变量都没有指定和设置，则默认从 config/config.dev.yaml文件中读取
项目使用 fsnotify https://github.com/fsnotify/fsnotify 来监听配置文件的变化，并在有变化时重新将结果反序列化到结构体中
```

### zap

```shell
日志初始化代码 gin-plus/pkg/logger/logger.go
日志切割归档代码 gin-plus/pkg/logger/logger.go
```