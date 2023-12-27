# mangodb

[![CN doc](https://img.shields.io/badge/文档-中文版-blue.svg)](mangodb.md)

- [组件说明](#组件说明)

---

### 组件说明
*  golang框架：
   * 原生mongo-driver
*  创建用户：
   * 原生mongo-driver
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
