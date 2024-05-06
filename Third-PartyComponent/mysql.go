package main

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// Props 道具
type Props struct {
	Id          int64  `gorm:"primaryKey"`
	PropType    int32  `gorm:"comment:道具类型;column:prop_type"`
	ServerId    int64  `gorm:"comment:服务器ID;column:server_id"`
	Name        string `gorm:"comment:道具名称;column:name"`
	Icon        string `gorm:"comment:道具图标;column:icon"`
	Description string `gorm:"comment:道具描述;column:description"`
	AddDesc     string `gorm:"comment:备注;column:add_desc"`
	CreatedAt   int64  `gorm:"autoCreateTime"`
	UpdatedAt   int64  `gorm:"autoUpdateTime"`
}

func mysqlInit() (connect *gorm.DB, err error) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(192.168.1.61:3306)/gamping?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		//Logger: func() logger.Interface {
		//	return logger.New(
		//		//将标准输出作为Writer
		//		log.New(os.Stdout, "\r\n[GormSlowSQL]", log.LstdFlags),
		//		logger.Config{
		//			//设定慢查询时间阈值为1ms
		//			SlowThreshold: 150 * time.Millisecond,
		//			//设置日志级别，只有Warn和Info级别会输出慢查询日志
		//			LogLevel: logger.Warn,
		//		},
		//	)
		//}(),
	})
	if err != nil {
		log.Printf("failed opening connection to mysql: %v", err)
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("failed get connection : %v", err)
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)           // 空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(50)           // 打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接可复用的最大时间
	err = sqlDB.Ping()
	if err != nil {
		log.Printf("failed ping connection : %v", err)
		return nil, err
	}
	return db, nil
}

func main() {
	connect, err := mysqlInit()
	if err != nil {
		return
	}
	var date []Props
	err = connect.WithContext(context.Background()).Where("prop_type = ?", 1).Find(&date).Error
	if err != nil {
		log.Printf("failed Find : %v", err)
		return
	}
	for _, value := range date {
		fmt.Println(value)
	}
}
