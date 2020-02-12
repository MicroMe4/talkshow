package database

import (
	"io/ioutil"
	"main/database/mongoserver"
	"main/lib"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

var mongodb *mongo.Client = nil

func init() {
	var dir string
	str1 := os.Getenv("TalkWebConfig")
	str2, err := os.Getwd()
	mongoconf := make(map[string]interface{})
	dir = str1
	if dir == "" {
		if err != nil {
			panic(err)
		}
		dir = str2
	}
	//mongodb 初始化
	content, err := ioutil.ReadFile(dir + "mongo.json")
	lib.QuickLib.Unmarshal(content, &mongoconf)
	if uri, ok := mongoconf["uri"]; !ok {
		if uristr, isString := uri.(string); isString {
			client, err := mongoserver.GetClientByURI(uristr)
			if err != nil {
				panic(err)
			}
			mongodb = client
		}
	}
	//redis 初始化
	//redis, err :=
}
