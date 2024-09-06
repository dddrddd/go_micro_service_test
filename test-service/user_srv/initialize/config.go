package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"test-service/user_srv/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug := GetEnvInfo("debug")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user_srv/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("user_srv/%s-debug.yaml", configFilePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.ServiceConfig); err != nil {
		panic(err)
	}
}
