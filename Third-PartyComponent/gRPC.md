# gRPC

[![CN doc](https://img.shields.io/badge/文档-中文版-blue.svg)](gRPC.md)

- [组件说明](#组件说明)
- [gRPC调用](#gRPC调用)
- 
---

### 组件说明
*  功能：
   * 底层采用HTTP2、支持双向流
   * 跨语言、跨平台
*  版本：
   * export ETCDCTL_API=3
*  安装：
   * etcd-server etcd-client
*  基本命令：
    * etcdctl get --prefix ""

### gRPC调用
*  服务间，在被调用的微服务中实现接口调用：
   * rpc ExecuteTaskTriggerEvent (ExecuteTaskTriggerEventReq) returns (EmptyReply) {}
*  对外，使用restful方法：
   * rpc GetUserTaskList (UserTaskListReq) returns (UserTaskListReply) {
     option (google.api.http) = {
     get: "/circle/v1/task/user_task_list"
     };
     }