package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type WhiteUser struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Phone   string             `bson:"phone"`
	SMSCode string             `bson:"sms_code"`
}

func mongoInit(ctx context.Context) (*mongo.Client, error) {
	mongodb, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:123456@192.168.1.61:27017"))
	if err != nil {
		panic(err)
	}
	err = mongodb.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	return mongodb, nil
}

func mongoOptions(ctx context.Context, connect *mongo.Client) error {
	var (
		filter = bson.M{"phone": "15947852365", "sms_code": "1234"}
	)
	count, err := connect.Database("gamping").Collection("login_white_users").CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	fmt.Printf("count of login_white_users: %d\n", count)
	return nil
}
