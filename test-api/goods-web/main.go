package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"test-api/goods-web/global"
	"test-api/goods-web/initialize"
	"test-api/goods-web/utils"
	"test-api/goods-web/utils/register/consul"
)

func main() {
	initialize.InitLogger()

	initialize.InitConfig()

	Router := initialize.Routers()

	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	initialize.InitSrvConn()

	viper.AutomaticEnv()
	debug := viper.GetBool("test_debug")
	if !debug {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}

	port := global.ServerConfig.Port
	//logger,_ := zap.NewProduction()
	//defer logger.Sync()
	//sugar := logger.Sugar()
	registerClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := registerClient.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败", err.Error())
	}
	zap.S().Debugf("启动服务器，端口：%d", port)
	if err := Router.Run(fmt.Sprintf(":%v", port)); err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}
}
