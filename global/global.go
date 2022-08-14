package global

import (
	ut "github.com/go-playground/universal-translator"
	proto "video/api/pb"
	"video/config"
)

var (
	Trans        ut.Translator // 表单验证
	ServerConfig *config.ServerConfig

	BarrageServerClient  proto.BarrageClient
	CategoryServerClient proto.CategoryClient
	EpisodesServerClient proto.EpisodesClient
	RegionServerClient   proto.RegionClient
	VideoServerClient    proto.VideoClient
	UserServerClient     proto.UserClient
)
