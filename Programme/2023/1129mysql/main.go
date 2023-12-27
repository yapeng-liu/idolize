package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Student struct {
	ID        int32
	Name      string
	UpdatedAt time.Time
	CreatedAt time.Time
}

type User struct {
	ID              int64
	Phone           string `gorm:"type:varchar(16);comment:手机号;index:idx_phone"`
	Nickname        string `gorm:"type:varchar(64);comment:昵称"`
	Gender          int32  `gorm:"type:tinyint(1);comment:性别 1：女；2：男"`
	Avatar          string `gorm:"type:varchar(512);comment:头像"`
	Background      string `gorm:"type:varchar(512);comment:背景"`
	Signature       string `gorm:"type:varchar(200);comment:签名"`
	Password        string `gorm:"type:varchar(256);comment:密码"`
	Salt            string `gorm:"type:varchar(4);comment:盐"`
	Suffix          int32  `gorm:"type:smallint(4);comment:后缀 1000~9999"`
	IsOnline        bool   `gorm:"default:false;comment:是否在线"`
	CreatedAt       int64  `gorm:"autoCreateTime"`
	UpdatedAt       int64  `gorm:"autoUpdateTime"`
	CreatedChannel  string `gorm:"type:varchar(20);comment:注册渠道"`
	LastLoginAt     int64  `gorm:"type:int(11);comment:最近登录时间"`
	State           int8   `gorm:"comment:用户账号状态 1：注销状态"`
	LastLoginMobile int8   `gorm:"comment:最近登录移动端类型 1：android，2：ios"`
	AppVersion      string `gorm:"comment:app版本"`
	IP              string `gorm:"comment:IP信息"`
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(192.168.1.61:3306)/gamping?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(150)
	sqlDB.SetConnMaxLifetime(time.Second * 25)
	var student []Student
	result := db.Debug().Find(&student, []int64{1, 3})

	fmt.Println(result.Error, "        ", result.RowsAffected, "        ", student)
	var ip string
	err = db.Debug().Model(User{}).Where("id", 220).Select("ip").Take(&ip).Error
	if err != nil {
		fmt.Println("err ", err)
	} else {
		fmt.Println("ip ", ip)
	}

}
