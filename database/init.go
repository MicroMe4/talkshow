package database

import (
	"io/ioutil"
	"main/lib"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

var mongodb *mongo.Client

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
	content, err := ioutil.ReadFile(dir + "mongo.json")
	lib.QuickLib.Unmarshal(content, &mongoconf)
	if uri, ok := mongoconf["uri"]; !ok {
		break
	} else {
		if uristr, isString := uri.(string); !isString {
			break
		} else {

		}

	}
}
