package api

import "go.mongodb.org/mongo-driver/bson/primitive"

//Key 可允许登陆到服务器的ID
type Key struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id"`
	AppKey string             `json:"appid" bson:"appid" xml:"appid"`
}
