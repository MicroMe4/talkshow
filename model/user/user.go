package user

import (
	"main/lib"
)
//User Without Salt.
type User struct {
	Username string `json:"username" xml:"username" bson:"username"`
	Password string `json:"password" xml:"password" bson:"password"`
}

func (u User) MarshalToJSON() ([]byte, error) {
	jsoniter.
}
