package api

import "go.mongodb.org/mongo-driver/bson/primitive"

//Key 登陆客户端控制系统
type Key struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id"`
	AppKey string             `json:"appid" bson:"appid" xml:"appid"`
}
