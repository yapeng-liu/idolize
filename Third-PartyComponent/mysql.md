# mysql

[![CN doc](https://img.shields.io/badge/文档-中文版-blue.svg)](mysql.md)

- [组件说明](#组件说明)

---

### 组件说明
*  golang框架：
   * gorm

### 使用
* 数据
~~~
数据库连接
db   *gorm.DB

数据库表
type Student struct {
   Id int64 `gorm:"id" json:"id"`
   Name  string   `gorm:"name" json:"name"`
}
~~~

*  根据主键检索单个对象
~~~
获取id=1记录（主键升序）
result := db.First(&Student,1)
result := db.First(&Student,"id=?",1)
// SELECT * FROM `students` WHERE `students`.`id` = 1 ORDER BY `students`.`id` LIMIT 1

获取id=1记录，没有指定排序字段
result := db.Take(&Student,1)
// SELECT * FROM `students` WHERE `students`.`id` = 1 LIMIT 1

获取id=1记录（主键降序）
result := db.Last(&Student,1)
// SELECT * FROM `students` WHERE `students`.`id` = 1 ORDER BY `students`.`id` DESC LIMIT 1



返回找到的记录数
result.RowsAffected
返回错误或者nil
result.Error

如果记录不存在 返回 ErrRecordNotFound 错误需要处理
errors.Is(result.Error, gorm.ErrRecordNotFound)

如何避免 ErrRecordNotFound 错误
result := db.Limit(1).Find(&Student)
~~~
*  根据主键检索多个对象
~~~
获取多条记录
db.Find(&users, []int{1,3})
// SELECT * FROM `students` WHERE `students`.`id` IN (1,3)


~~~

