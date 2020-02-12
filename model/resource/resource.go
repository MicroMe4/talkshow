package resource

//Resource 资源文件
type Resource struct {
	ResourceLocation string `json:"audio" xml:"audio" bson:"audio"`
	Sender           string `json:"owner" xml:"owner" bson:"owner"`
	Receiver         string `json:"receiver" xml:"receiver" bson:"receiver"`
}
