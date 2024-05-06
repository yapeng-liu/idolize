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

##### .deb文件安装
~~~
sudo apt install path_to_deb_file

sudo dpkg -i path_to_deb_file
~~~