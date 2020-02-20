package database

import (
	"io/ioutil"
	"os"

	"github.com/MicroMe4/talkshow/database/mongoserver"
	"github.com/MicroMe4/talkshow/lib"

	"github.com/go-redis/redis/v7"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongodb *mongo.Client = nil
var redisdb *redis.Client = nil

func init() {
	var dir string
	str1 := os.Getenv("TalkWebConfig")
	str2, err := os.Getwd()
	conf := make(map[string]interface{})
	dir = str1
	if dir == "" {
		if err != nil {
			panic(err)
		}
		dir = str2
	}
	//mongodb 初始化
	content, err := ioutil.ReadFile(dir + "mongo.json")
	lib.QuickLib.Unmarshal(content, &conf)
	if uri, ok := conf["uri"]; !ok {
		if uristr, isString := uri.(string); isString {
			client, err := mongoserver.GetClientByURI(uristr)
			if err != nil {
				panic(err)
			}
			mongodb = client
		}
	}
	content, err = ioutil.ReadFile(dir + "redis.json")
	lib.QuickLib.Unmarshal(content, &conf)
	redisdb = redis.NewClient(&redis.Options{
		Addr:     conf["address"].(string),
		Password: conf["password"].(string),
	})
	if redisdb == nil {
		panic("failed assigning redis memory database")
	}
}
