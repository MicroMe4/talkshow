package user

import (
	matchutil "crypto/sha512"
	"encoding/hex"

	"github.com/MicroMe4/talkshow/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User Structure.
type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Username string             `json:"username" xml:"username" bson:"username"`
	Password string             `json:"password" xml:"password" bson:"password"`
	Salt     []byte             `json:"salt" xml:"salt" bson:"salt"`
}

//MarshalToJSON 转换为JSON
func (u User) MarshalToJSON() ([]byte, error) {
	return lib.QuickLib.Marshal(u)
}

//MarshalToBSON 转换为BSON
func (u User) MarshalToBSON() ([]byte, error) {
	return bson.Marshal(u)
}

//TryLogin 登陆函数
func (u User) TryLogin(pwd string) bool {
	code := matchutil.New()
	code.Write([]byte(pwd))
	return hex.EncodeToString(code.Sum(u.Salt)) == u.Password
}
