# mysql

[![CN doc](https://img.shields.io/badge/文档-中文版-blue.svg)](mysql.md)

- [组件说明](#组件说明)

---
###  GUI
* Ubuntu系统-Workbench
  * sudo apt-get install mysql-workbench-community

### 组件说明
*  golang框架：
  * gorm

### 使用
* 官方文档
  *  https://gorm.io/zh_CN/
* 索引创建
  * 跟随建表一起创建
  * 建完表后 单独创建
* 索引生效
  * 对于表示状态字段-status增加索引
    * 如果where条件只包括 状态索引字段，查询效率有明显提升
    * 如果where条件包括 状态索引字段+非索引字段 状态值的数据量较少时 ，查询效率有提升
    * 如果where条件包括 状态索引字段+非索引字段 状态值的数据量较多时 ，查询效率会降低























