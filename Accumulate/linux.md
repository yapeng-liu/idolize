# linux

[![CN doc](https://img.shields.io/badge/文档-中文版-blue.svg)](linux.md)

- [命令](#命令)

---
### 命令
##### 压缩与解压缩
~~~
压缩test文件
tar -zcvf test.tar /usr/test
解压
tar -zxvf test.tar
~~~

##### 查看文件占用大小
~~~
du -h --max-depth=0 *
~~~

##### .deb文件安装与删除
~~~
sudo apt install path_to_deb_file

sudo dpkg -i path_to_deb_file

dpkg -r todesk

dpkg -l todesk 

killall snap-store

sudo snap refresh snap-store
~~~

##### 系统架构
~~~
lscpu
x86_64:表示基于 Intel 或 AMD 的 64 位架构
ARM:

~~~