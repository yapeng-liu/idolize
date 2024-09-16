# mangodb

[![CN doc](https://img.shields.io/badge/文档-中文版-blue.svg)](mongo)

- [组件说明](#组件说明)
- [备份与恢复](#备份与恢复)

---

### 组件说明
*  golang框架：
   * 原生mongo-driver
*  GUI工具：
   * MongoDB Compass
### docker-compose文件
~~~
version: "3"

networks:
  gamping-test:

services:
  mongo6:
    image: mongo:6.0
    networks:
      - gamping-test
    ports:
      - "27018:27017"
    volumes:
      - ./data/mongo6/data:/mongo/data
      - ./data/mongo6/log:/mongo/log
      - ./data/mongo6/mongo.keyfile:/mongo/mongo.keyfile
    command: "mongod --dbpath /mongo/data --logpath /mongo/log/mongo.log --replSet preSet --keyFile /mongo/mongo.keyfile --oplogSize 128"
    environment:
      - MONGO_INITDB_ROOT_USERNAME=admin
      - MONGO_INITDB_ROOT_PASSWORD=123456
    entrypoint:
      - bash
      - -c
      - |
        chmod 400 /mongo/mongo.keyfile
        chown 999:999 /mongo/mongo.keyfile
        exec docker-entrypoint.sh $$@
~~~
### 备份与恢复
~~~
//启动mongo,使用指定path存储数据
mongod --dbpath [path]
//进入mongo命令行
mongo
//选择数据库
use admin
//插入用户数据
db.createUser({user:"admin",pwd:"123456",roles:["root"]})
db.createUser({user: "random",pwd: "aWWKyAdDoU2Yi5doL66c",roles:[{role: "readAnyDatabase", db: "admin"},{role: "read", db: "local"}]})

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

docker exec gamping_mongo-3_1 mongodump -u admin -p 123456 -o /data/dumpFile

docker cp -a gamping_mongo-3_1:/data/dumpFile ./mongo1

rs.initiate({
_id: "preSet",
members: [
{_id: 0, host: "192.168.1.61:27018"}
]
});

cfg.members[0].host = "172.16.16.15:27017"

var cfg = rs.conf();
cfg.members.forEach(function(member) {
if (member._id == 0) {
member.host = "114.132.183.244:27017";
}
});
