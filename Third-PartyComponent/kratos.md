# kratos

[![CN doc](https://img.shields.io/badge/文档-中文版-blue.svg)](kratos.md)

- [配置文件](#配置文件)
- [服务创建启动](#服务创建启动)

---
### 配置文件
*  文件名：conf.proto
*  命令：make config
*  生成：conf.pb.go

### 服务创建启动
*  协议接口：
   * api/helloWorld/v1/student.proto
*  客户端交互：
   * kratos proto client  api/helloWorld/v1/student.proto
*  服务端功能实现：
   * internal层（权限校验、参数校验）
     * kratos proto server  api/helloWorld/v1/student.proto -t internal/service
   * biz层（逻辑处理）
     * 可以将grpc中的传参解析之后在biz层处理
     * 自己创建文件及函数，需要注入
   * data层 （数据库及其它存储通信操作）
     * 数据层的具体操作，要能被复用
     * 自己创建文件及函数，需要注入


## 项目
* layout 模板
* .dev 文档
* aegis
* beer-shop
* gateway


## 工程化最佳实践

### 配置文件
* 
~~~
    1.默认本地文件数据源、xx.yaml,proto文件
        1.统一模板配置
        2.配置校验
        3.管理配置
        4.跨语言、跨平台
       
    2.自定义数据源，实现接口 
        type KeyValue struct{
            Key string
            Value []byte
            Format string
        }
        type Source interface{
            Load()([]*KeyValue,error)
            Watch()(Watch,error)
        }
        type Watcher interface{
            Next()([]*KeyValue,error)
            Stop()error
        }
    统一支持：
        热更新
        数据源解码
        配置填充、修改配置文件
~~~

### Error Handling
~~~
    message Error{
        int32 code = 1; //错误码
        string reason = 2; //错误原因 Code不足可以使用这个
        string message = 3; //错误信息
        map<string,string>metadata = 4; //错误元信息
    }
    不要将服务错误传播到客户端
        举个例子，你现在要跟移动端说我有一个接口，那么这个接口会返回哪些错误码，你始终讲不清楚，你为什么讲不清楚呢？
        因为我们整个微服务的调用链是 A 调 B，B 调 C，C 调 D，D 的错误码会一层层透传到 A，那么 A 的错误码可能会是 ABCD 错误码的并集，
        你觉得你能描述出来它返回了哪些错误码吗？根本描述不出来
    建议：
        1.隐藏实现详细信息、敏感信息
        2.调整负责该错误的一方，如：从一个服务接收invalid_argument错误的服务应该将包装成internal或其它的传播给调用者
        
        比如你返回的错误码是4，代表商品已下架，我对这个错误很感兴趣，但是错误码4 在我的项目里面已经被用了，
        我就把它翻译为我还没使用的错误码6，这样每次翻译的时候就可以对上一层你的调用者，你就可以交代清楚你会返回错误码，因为都是你定义的，而且是你翻译的，
        你感兴趣的才翻译，你不感兴趣的通通返回 500 错误，就是内部错误，或者说 unknown，就是未知错误，这样你每个 API 都能讲清楚自己会返回哪些错误码
~~~

### Layout
* 统一仓库、统一包名 ？proto文件如何关联的服务  
* protobuf管理 buf


### 事务管理
* 业务层事务、存储层事务？到底放在哪里
  * 事务逻辑作用于用例层（业务逻辑），而不在持久层
  * 将业务逻辑于事务代码分开，需要添加事务时进行简单封装
* 更多实现在kratos仓库的examples目录的transaction
如下举例：
~~~
新模式：
    1.biz层抽象
    useCase层通过wire由data注入tx Transaction接口
    
    func NewUserUsecase(repo UserRepo, followRepo FollowRepo, logger log.Logger, tx Transaction) *UserUsecase {
	    return &UserUsecase{repo: repo, followRepo: followRepo, log: log.NewHelper(logger), tx: tx}
	}
    
    type Transaction interface {
	    ExecTx(context.Context, func(ctx context.Context) error) error
    }
    
    2.data层实现
    
    func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
        return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
            ctx = context.WithValue(ctx, contextTxKey{}, tx)
            return fn(ctx)
        })
    }
    
    // DB 根据此方法来判断当前的 db 是不是使用 事务的 DB
    func (d *Data) DB(ctx context.Context) *gorm.DB {
        tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
        if ok {
            return tx
        }
        return d.db.WithContext(ctx)
    }
    

    // NewTransaction .
    func NewTransaction(d *Data) biz.Transaction {
	    return d
    }
~~~

## 微服务治理

* 单体架构
  * 维护、开发等问题
* 面向服务架构
  * 化繁为简 分而治之

#### 服务发现及负载均衡
~~~
import  "github.com/go-kratos/kratos/v2/selector/wrr"
import  "github.com/go-kratos/kratos/v2/selector/filter"

// 创建路由 Filter：筛选版本号为"2.0.0"的实例
filter :=  filter.Version("2.0.0")
// 创建 P2C 负载均衡算法 Selector，并将路由 Filter 注入
selector.SetGlobalSelector(wrr.NewBuilder())

hConn, err := http.NewClient(
  context.Background(),
  http.WithEndpoint("discovery:///helloworld"),
  http.WithDiscovery(r),
  http.WithNodeFilter(filter)
)

~~~


#### 限流、熔断等可用性算法
~~~
var opts = []http.ServerOption{
    http.Middleware(
        // 默认 bbr limiter
        ratelimit.Server(),
        // 自定义 limiter
        //ratelimit.Server(ratelimit.WithLimiter(limiter)),
    ),
}

srv := http.NewServer(opts...)

~~~

#### 服务可观测性
~~~

package server

import (
    "github.com/go-kratos/kratos/v2/middleware/tracing"
    "github.com/go-kratos/kratos/v2/transport/grpc"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/sdk/resource"
    tracesdk "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// 设置全局trace
func initTracer(url string) error {
    // 创建 Jaeger exporter
    exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
    if err != nil {
        return err
    }
    tp := tracesdk.NewTracerProvider(
        // 将基于父span的采样率设置为100%
        tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
        // 始终确保在生产中批量处理
        tracesdk.WithBatcher(exp),
        // 在资源中记录有关此应用程序的信息
        tracesdk.WithResource(resource.NewSchemaless(
            semconv.ServiceNameKey.String("kratos-trace"),
            attribute.String("exporter", "jaeger"),
            attribute.Float64("float", 312.23),
        )),
    )
    otel.SetTracerProvider(tp)
    return nil
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, executor *service.ExecutorService) *grpc.Server {
    err := initTracer("http://localhost:14268/api/traces")
    if err != nil {
        panic(err)
    }
    //tr := otel.Tracer("component-main")
    var opts = []grpc.ServerOption{
        grpc.Middleware(
            tracing.Server(),
        ),
    }
    // ...
}

~~~

#### 动态路由

将北京的请求打到北京的集群

标识请求到特定的服务


## API Gateway网关设计与最佳实践
* 第一次重构-SOA
  * 微服务直接调用，强耦合
  * 客户端多次请求、客户端聚合数据
  * 各协议不统一
  * 统一逻辑无法收敛、如安全认证、限流
* 网关演进-BFF 1.0
  * 新增app-interface 统一出口
  * 差异服务、数据裁减聚合、针对终端定制化API
  * 动态升级、升级服务
* BFF 2.0
  * 多个BFF
* API Gateway
  * 将路由、认证、限流、安全全部上移 
  * 目标
    * API元信息
    * 流量调度
    * 隔离保护
    * 访问控制
    * 可观测性

## 数据安全
* 数据加密
  * 数据加解密
    * AES-256
      * 加密：明文数据->标记位+版本信息+密文->Base64编码->存储
      * 解密：相反
      * 查询：Scan\Value自动加解密
    * TDES和TDEA
    * RSA
* 数据脱敏
  * 自定义log
  * 脱敏中间件
  * 脱敏算法
    * 遮盖脱敏
    * 哈希脱敏
    * 替换脱敏
      * 字符替换
## 运维数仓
* CMDB3

## 架构演进
* 发展历程
  * 从零到一
  * 微服务化
  * 容器化
    * CPU调度模式选择
      * CPUSET，绑定CPU
      * CFS，时间片分配
    * PHP服务问题：
      * 进程数过多
      * CGROUP泄漏
      * CPU Burst
    * 突发流量
      * 扩容
      * 限流
      * 流量隔离
  * PHP到Golang
    * 并发
  * 网关架构
    * API聚合
* 热点专题讨论
  * 热点数据
    * CDN\redis\内存等提前缓存
  * 流量放大
    * 接口返回不必要的字段的问题
      * GraphQL 字段选择
      * GRPC FieldMask
    * 服务请求多次同样的数据
      * 断开非必要的依赖、对于内部服务由调用方通过参数传递依赖
  * 活动保障
    * 场景分析
    * 容量评估
    * 活动预案
      * 接口超时
      * 缓存故障
      * 数据库故障
      * 流量超限
      * 磁盘写满
    * 现场保障