package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"test-api/user-web/global"
)

// 可以获取环境变量
func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug := GetEnvInfo("test_debug")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user-web/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("user-web/%s-debug.yaml", configFilePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
	// 动态监视配置文件是否变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("config file changed: %s", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
		zap.S().Infof("配置信息: %v", global.ServerConfig)
	})
}
