# Set

[![CN doc](https://img.shields.io/badge/文档-中文版-blue.svg)](set.md)

- [远程服务器免密登录](#远程服务器免密登录)

---

### 远程服务器免密登录
*  时间：20230925
*  功能：免密登录、复制文件
*  本地生成密钥，已经生成过则不需要：
   * ssh-keygen -t rsa
*  设置远程服务器密钥验证：
   * scp ~/.ssh/id_rsa.pub user@remote.server.com:~/.ssh/authorized_keys
   * scp ~/.ssh/id_rsa.pub root@159.75.92.101:~/.ssh/authorized_keys