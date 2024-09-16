# redis

[![CN doc](https://img.shields.io/badge/文档-中文版-blue.svg)](redis.md)

- [组件说明](#组件说明)

---

### 组件说明
*  golang框架：
   * go-redis

### docker-compose文件
~~~
version: "3"

networks:
  gamping-test:

services:
  redis6:
    image: 'redis:6.2'
    ports:
      - 6379:6379
    volumes:
      - ./data/redis6/data:/data
    command: redis-server --requirepass LJFDSLjl89234 --logfile redis.log
    networks:
      - gamping-test
~~~

### 数据备份与恢复
* 本次为docker5.0 升级到 6.2 做数据备份与恢复
* 登录5.0实例备份,备份文件名默认为dump.rdb，文件保存的路径为挂载的data中 
   ~~~
   redis 127.0.0.1:6379> SAVE 
   OK
   ~~~
* 恢复
  * 登录6.2实例，查看数据保存目录
  ~~~
  redis 127.0.0.1:6379> CONFIG GET dir
  1) "dir"
  2) "./data"
  ~~~
  * 将备份文件(dump.rdb)移动到6.0实例中挂载的data下即可
  * 重启6.2实例
* 恢复注意
  * 重启6.0实例发现dump.rdb文件被覆盖
  * 第一点: 关闭 appendonly
    * appendonly no
  * 第二点: 容器启动顺序
    * 不能直接使用 docker restart 6.2实例
    * docker stop时会备份 覆盖本地dump.rdb
  * 以下为正确的数据恢复流程
    * docker stop 6.2实例
    * 将备份文件(dump.rdb)移动到6.2实例中挂载的data下
    * docker start 6.2实例