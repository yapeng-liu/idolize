# mangodb

[![CN doc](https://img.shields.io/badge/文档-中文版-blue.svg)](mangodb.md)

- [组件说明](#组件说明)
- [备份](#备份)

---

### 组件说明
*  golang框架：
   * 原生mongo-driver
*  GUI工具：
   * MongoDB Compass
### 备份ls
~~~
//启动mongo,使用指定path存储数据
mongod --dbpath [path]
//进入mongo命令行
mongo
//选择数据库
use admin
//插入用户数据
db.createUser({user:"admin",pwd:"123456",roles:["root"]})
//查看结果
show users
//恢复备份数据
mongorestore ./mongo3-2023-10-19-11-44-19/
~~~
### 建立索引
~~~
 超时错误
    connection(172.17.0.1:30005[-33]) incomplete read of message header: context deadline exceeded latency=3.001131637
 查看索引
    db.'表名'.getIndexes()
 建立索引
    db.activity.createIndex({"hot":-1});
    db.comment.createIndex({"topic_id":1});
    db.like.createIndex({"object_id":1});
 删除索引
    db.activity.dropIndex("hot_-1")
~~~

