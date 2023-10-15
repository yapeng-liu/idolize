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
