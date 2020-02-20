package resource

import (
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Resource 资源数据
type Resource struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id" xml:"_id"`                      //ID 数据唯一性标志
	DescriptorStr string             `json:"descriptor" xml:"descriptor" bson:"descriptor"` //DescriptorStr 描述信息 （A:B:C:D） A为数据类型(Audio, Video, Text...),B为编码方式加解密(BASE64/BASE32/AES...)，支持多层，C为D代表的意义，即如何解析它，D为数据本身
	Sender        string             `json:"owner" xml:"owner" bson:"owner"`                //Sender 发送者
}

//Descriptor 资源描述符
type Descriptor struct {
	types       uint64
	code        uint64
	source      uint64
	description string
}

const (
	//TextType 文档类型
	TextType uint = 1 << iota
	//PictureType 图片类型
	PictureType
	//StreamingType 流数据（中转）类型
	StreamingType
)

const (
	//Base64Encoding Base64编码
	Base64Encoding uint = 1 << iota
	//AESEncoding AES加密编码
	AESEncoding
	//Plain 无编码
	Plain
)

const (
	//FileSystem 文件系统信息（通过文件系统获取信息）
	FileSystem uint = 1 << iota
	//URI URL信息
	URI
	//PlainContent 包含在数据内的信息
	PlainContent
)

//DescriptorMustParse 文件描述符分离，无法分离时将会Panic
func DescriptorMustParse(rdstring string) *Descriptor {
	val := strings.Split(rdstring, ":")
	if len(val) != 4 {
		panic("descriptor invalid")
	}
	key, err0 := strconv.ParseUint(val[0], 10, 0)
	codec, err1 := strconv.ParseUint(val[1], 10, 0)
	contentlvl, err2 := strconv.ParseUint(val[2], 10, 0)
	if err0 != nil || err1 != nil || err2 != nil {
		panic("atoi failed")
	}
	return &Descriptor{
		types:       key,
		code:        codec,
		source:      contentlvl,
		description: val[3],
	}
}

//GetType 获取类型
func (d Descriptor) GetType() uint64 {
	return d.types
}

//GetDescription 获取描述
func (d Descriptor) GetDescription() string {
	return d.description
}

// GetSource 获取信息来源
func (d Descriptor) GetSource() uint64 {
	return d.source
}

// GetCodec 获取编码信息
func (d Descriptor) GetCodec() uint64 {
	return d.code
}

//BeforeInsert 加入数据库之前用的函数
func (r Resource) BeforeInsert() {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	DescriptorMustParse(r.DescriptorStr)
}

//TableName 存储时候的表名/集合名称
func (r Resource) TableName() string {
	return "resource"
}
