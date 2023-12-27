# docker

[![CN doc](https://img.shields.io/badge/文档-中文版-blue.svg)](docker.md)

- [组件说明](#组件说明)

---

### 组件说明
* 在一个已经exit的docker容器中修改配置文件
~~~
//列举容器具体信息
docker inspect [CONTAINER ID]
//定位XXXXX信息
 "GraphDriver":{
   "Data": {
     XXXXXXXX
   }
 }
 //打开容器信息文件进行相应修改
~~~
* 磁盘清理
~~~
docker system prune

docker system prune -a 

存在僵尸文件 第二天查看磁盘空间
~~~
* 修改配置文件重启docker compose容器
~~~
docker-compose stop worker       // go to hibernate
docker-compose rm worker        // shutdown the PC 
docker-compose create worker     // create the container from image and put it in hibernate
docker-compose start worker //bring container to life from hibernation

~~~