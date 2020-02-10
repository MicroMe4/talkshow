package audio

//Audio 声音录像资源
type Audio struct {
	Audioloc string `json:"audio" xml:"audio" bson:"audio"`
	Owner    string `json:"owner" xml:"owner" bson:"owner"`
	CreateAt string `json:"create" xml:"create" bson:"create"`
}

func (a *Audio) AssignOwner() {

}
