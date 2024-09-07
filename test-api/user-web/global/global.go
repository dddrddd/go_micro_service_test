package global

import (
	ut "github.com/go-playground/universal-translator"
	"test-api/user-web/config"
	"test-api/user-web/proto"
)

// 需要把配置信息设置为全局变量
var (
	ServerConfig *config.ServerConfig
	Trans        ut.Translator
	NacosConfig  *config.NacosConfig

	UserSrvClient proto.UserClient
)
