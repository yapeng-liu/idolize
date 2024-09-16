package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type WhiteUser struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Phone   string             `bson:"phone"`
	SMSCode string             `bson:"sms_code"`
}

type Activity struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	UserId         int64              `bson:"user_id"`
	ServerId       int64              `bson:"server_id"`
	ChannelId      int64              `bson:"channel_id"`
	Title          string             `bson:"title"`
	Content        string             `bson:"content"`
	Pictures       []string           `bson:"pictures"`
	Video          string             `bson:"video"`
	VideoUuid      string             `bson:"video_uuid"`
	VideoThumbnail string             `bson:"video_thumbnail"`
	LikeCount      int32              `bson:"like_count"`
	CommentCount   int32              `bson:"comment_count"`
	CollectCount   int32              `bson:"collect_count"`
	//ShareCount     int32              `bson:"share_count"`
	BrowseCount int32  `bson:"browse_count"`
	SeeType     int8   `bson:"see_type"`
	Attrs       int32  `bson:"attrs"`
	ExtendType  int8   `bson:"extend_type"`
	ExtendId    string `bson:"extend_id"`
	Status      int32  `bson:"status"`
	Visible     int32  `bson:"visible"`
	//Hot             float64              `bson:"hot"`
	//HotDecay        float64              `bson:"hot_decay"` //热度衰减值
	CreatedAt       int64    `bson:"created_at"`
	UpdatedAt       int64    `bson:"updated_at"`
	UpdatedText     string   `bson:"updated_text"`
	ContentLink     string   `bson:"content_link"`
	Essence         bool     `bson:"essence"`
	IsLongEssay     bool     `bson:"is_long_essay"`
	FavoritesCount  int32    `bson:"favorites_count"`
	TopicIds        []uint64 `bson:"topic_ids"`
	CatalogId       []int64  `bson:"catalog_id"`
	PublishType     int32    `bson:"publish_type"`
	ActivityDisplay int32    `bson:"activity_display"`
	AskCount        int32    `bson:"ask_count"`
	TagId           []int64  `bson:"tag_id"`
	IsRecommend     bool     `bson:"is_recommend"`
	PublishTime     int64    `bson:"publish_time"` //定时发布时间
	HotAt           int64    `bson:"hot_at"`
	EditedAt        int64    `bson:"edited_at"`
	VoteMode        int32    `bson:"vote_mode"`       //投票方式
	StartTime       int64    `bson:"start_time"`      //开始时间
	EndTime         int64    `bson:"end_time"`        //结束时间
	UserVoteCount   int64    `bson:"user_vote_count"` //用户参与投票数量
}

func mongoInit(ctx context.Context) (*mongo.Client, error) {
	opts := options.Client().ApplyURI("mongodb://admin:74dZB3G3gUHJw3NRqwT9@192.168.30.172:27010,192.168.30.172:27011,192.168.30.172:27012")
	//opts := options.Client().ApplyURI("mongodb://admin:123456@192.168.1.61:27017,192.168.1.61:27018,192.168.1.61:27019").SetReadPreference(readpref.Nearest())
	mongodb, err := mongo.Connect(ctx, opts)
	if err != nil {
		fmt.Println("Connect err", err)
		panic(err)
	}
	err = mongodb.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Ping err", err)
		panic(err)
	}
	fmt.Println("mongodb connect success")
	return mongodb, nil
}

func mongoOptions(ctx context.Context, connect *mongo.Client) error {
	for i := 20; i < 30; i++ {
		activity := &Activity{
			UserId:    int64(i),
			ServerId:  200 + int64(i),
			ChannelId: 300 + int64(i),
			Title:     "好好好",
			Content:   "12313213",
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
		count, err := connect.Database("feed").Collection("activity").InsertOne(ctx, activity)
		if err != nil {
			fmt.Println("InsertOne err", err)
			return err
		}
		fmt.Printf("ID = %s\n", count.InsertedID.(primitive.ObjectID).Hex())
	}
	return nil
}

//db.AaA.insertOne({"phone": "15947852365", "sms_code": "1234"})
