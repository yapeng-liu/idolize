package main

import (
	"context"
	"database/sql"
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
	Status          int32  `gorm:"comment:状态"`
	SetStatusAt     int64  `gorm:"comment:设置状态时间"`
	Idea            string `gorm:"comment:想法"`
	OnlineAt        int64  `gorm:"comment:上线时间"`
	OfflineAt       int64  `gorm:"comment:下线时间"`
}

type ServerPo struct {
	ID                  int64
	UserID              int64        `gorm:"comment:用户id"`
	Type                int32        `gorm:"comment:服务器类型  0: 个人 1:工会 2: 官方 3:特殊圈子"`
	Name                string       `gorm:"size:128;comment:服务器名称"`
	Avatar              string       `gorm:"size:256;comment:服务器头像"`
	Overview            string       `gorm:"size:512;comment:服务器概述"`
	Count               int32        `gorm:"comment:服务器人数"`
	Show                sql.NullBool `gorm:"default:true;comment:服务器是否公开"`
	Background          string       `gorm:"size:256;comment:背景图片"`
	Status              int32        `gorm:"comment:服务器状态0：正常 1：已删除"`
	TopActivityID       string       `gorm:"size:256;comment:置顶动态id;not null"`
	TopActivityTime     int64        `gorm:"comment:置顶动态时间;not null"`
	ActivityNum         int32        `gorm:"comment:动态数;not null"`
	LastActivityTime    int64        `gorm:"comment:最后一条动态时间;not null"`
	ActivityChannelName string       `gorm:"comment:动态频道名;not null"`
	HasGame             int32        `gorm:"comment:是否有游戏"`
	GameType            int32        `gorm:"comment:服务器所对应的游戏类型"`
	BindGameType        int8         `gorm:"comment:绑定游戏方式：1返回所有大区，2返回有角色的大区"`
	BindGameStatus      int8         `gorm:"comment:强制绑定游戏账号状态，按二进制位控制开关，0001：商城兑换强制，0010：见面礼强制，0100：加入服务器强制 1000邮件绑定 10000抽奖强制绑定"`
	UnbindGameCdTime    int64        `gorm:"comment:服务器所对应的游戏类型解绑后冷却时间"`
	TodayAtTimes        int32        `gorm:"today_at_times" json:"today_at_times"` // 今日at全部成员的次数
	LastAtTime          int64        `gorm:"last_at_time" json:"last_at_time"`     // 上一次at全部成员的时间
	ActivityIcon        string       `gorm:"comment:频道tab页活动icon"`
	CreatedAt           int64        `gorm:"autoCreateTime"`
	UpdatedAt           int64        `gorm:"autoUpdateTime"`
	FirstIcon           string       `gorm:"first_icon" json:"first_icon"`                 //首页背景
	GuideType           int32        `gorm:"guide_type" json:"guide_type"`                 //新手引导类型
	FindFriendIcon      string       `gorm:"find_friend_icon" json:"find_friend_icon"`     //找好友浮标
	IsPopDayWelfare     int32        `gorm:"is_pop_day_welfare" json:"is_pop_day_welfare"` //是否弹出天天领福利
}

func (cm *ServerPo) TableName() string {
	return "servers"
}

type ActivityTag struct {
	gorm.Model
	Tag               string `gorm:"column:tag"`
	Des               string `gorm:"column:des"`
	BindActivityCount int64  `gorm:"column:bind_activity_count"`
}

func (ActivityTag) TableName() string {
	return "activity_tag"
}

func mysqlInit(ctx context.Context) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(192.168.1.61:3306)/gamping?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		//db, err := gorm.Open(mysql.Open("root:9bWVj.9jSD@tcp(42.194.237.166:30001)/gamping?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
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

func mysqlOptions(ctx context.Context, connect *gorm.DB) error {
	//var date []Props
	//err := connect.WithContext(ctx).Where("prop_type = ?", 1).Limit(1).Find(&date).Error
	//if err != nil {
	//	log.Printf("failed Find : %v", err)
	//	return err
	//}
	//for _, value := range date {
	//	fmt.Println(value)
	//}
	//bT := time.Now()
	//var user User
	//err := connect.WithContext(ctx).Where("phone = ?", "13756249027").Order("state").First(&user).Error
	//if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	//	return err
	//}
	//eT := time.Since(bT)
	//fmt.Println("Run time: ", eT)
	//var circles []ServerPo
	//fmt.Println(len(circles))
	//err := connect.WithContext(ctx).Model(&ServerPo{}).Where("status = ?", 0).Find(&circles).Error
	//if err != nil {
	//	return err
	//}
	//fmt.Println(len(circles))
	//list := make([]*Circle, 0, len(circles))
	//for _, c := range circles {
	//	list = append(list, serverPo2Biz(c))
	//}
	//
	//fmt.Println(len(list))

	if err := connect.Model(&ActivityTag{}).Where("id in ?", []int64{5, 6}).Update("bind_activity_count", gorm.Expr("bind_activity_count + ?", 1)).Error; err != nil {
		fmt.Println("aaa", err)
		return err
	}

	return nil
}

func serverPo2Biz(c ServerPo) *Circle {
	hasGame := false
	isPopDayWelfare := false
	if c.HasGame > 0 {
		hasGame = true
	}
	if c.IsPopDayWelfare > 0 {
		isPopDayWelfare = true
	}
	return &Circle{
		ID:                  c.ID,
		UserID:              c.UserID,
		Type:                c.Type,
		Name:                c.Name,
		Avatar:              c.Avatar,
		Overview:            c.Overview,
		Count:               c.Count,
		Show:                c.Show.Bool,
		Background:          c.Background,
		Status:              c.Status,
		TopActivityID:       c.TopActivityID,
		TopActivityTime:     c.TopActivityTime,
		ActivityNum:         c.ActivityNum,
		LastActivityTime:    c.LastActivityTime,
		ActivityChannelName: c.ActivityChannelName,
		HasGame:             hasGame,
		GameType:            c.GameType,
		BindGameType:        c.BindGameType,
		BindGameStatus:      c.BindGameStatus,
		UnbindGameCdTime:    c.UnbindGameCdTime,
		ActivityIcon:        c.ActivityIcon,
		CreatedAt:           c.CreatedAt,
		UpdatedAt:           c.UpdatedAt,
		FirstIcon:           c.FirstIcon,
		GuideType:           c.GuideType,
		FindFriendIcon:      c.FindFriendIcon,
		IsPopDayWelfare:     isPopDayWelfare,
	}
}

type Circle struct {
	ID                  int64
	UserID              int64
	Type                int32
	Name                string
	Avatar              string
	Overview            string
	Count               int32
	Online              int32
	Show                bool
	Background          string
	Status              int32
	TopActivityID       string
	TopActivityTime     int64
	ActivityNum         int32
	LastActivityTime    int64
	ActivityChannelName string
	HasGame             bool
	GameType            int32
	BindGameType        int8
	BindGameStatus      int8
	UnbindGameCdTime    int64
	ActivityIcon        string
	CreatedAt           int64
	UpdatedAt           int64
	FirstIcon           string
	GuideType           int32
	FindFriendIcon      string
	IsPopDayWelfare     bool
}
