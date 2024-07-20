# git

[![CN doc](https://img.shields.io/badge/文档-中文版-blue.svg)](git.md)

- [组件说明](#组件说明)

---

### 组件说明
💎️分支提交修改
 1.1 创建并新建feature分支
       git checkout -b feature
 1.2 分支已存在则切换分支
       git checkout feature
 2 拉取最新代码
       git pull --rebase origin develop
 3 基于自己的分支修改代码
 4 公共分支修改，先将自己的代码改动进行缓存
       git stash
 5 拉取更新最新代码
       git pull --rebase origin develop
 6 将自己缓存的代码释放出来
       git stash pop
 7 如果存在冲突则解决冲突
 8 添加自己修改部分
       git add .
 9 增加提交说明
       git commit -m "fix:fix bug"
 10 上传
      git push origin feature
 11.1.1 请求合并
      页面发起 merge request
 11.1.2 分支合并
      页面进行合并 merge
 11.2.1 OR 切换到develop分支
      git checkout develop
 11.2.2 分支合并
      git merge feature
 11.2.3 上传到远程分支
      git push 