package lib

import (
	"encoding/base64"

	jsoniter "github.com/json-iterator/go"
)

//QuickLib Json编码库
var QuickLib jsoniter.API = jsoniter.ConfigFastest

//DecodeBase Base编码器
var DecodeBase = base64.StdEncoding

