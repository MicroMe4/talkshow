package seeker

//VersionCode 版本号
type VersionCode = uint8

//Seeker 登陆记录设定
type Seeker struct {
	Username    string      `json:"username" xml:"username"`
	LogLocation string      `json:"logloc" xml:"logloc"`
	IPVersion   VersionCode `json:"type" xml:"type"`
}

const (
	//IPv4 IPv4版本号
	IPv4 VersionCode = 4
	//IPv6 IPv6版本号
	IPv6 VersionCode = 6
)
