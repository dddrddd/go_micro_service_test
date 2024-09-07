package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"test-api/user-web/global"
	"test-api/user-web/proto"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%s/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserInfo.Name), //这一部分使用了"github.com/mbobakov/grpc-consul-resolver"，固定格式，标签可以参考官方文档
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), //这部分也是
	)
	if err != nil {
		zap.S().Fatalf("用户服务失败")
	}
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}

func InitSrvConn2() { //这是没在使用负载均衡前的代码
	//从注册中心获取用户服务的信息
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	userSrvHost := ""
	userSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserInfo.Name)) //拿着服务字段过滤，有很多字段可选择
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}
	if userSrvHost == "" {
		zap.S().Fatal("连接用户服务失败")
		return
	}

	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList]连接用户服务失败",
			"msg", err.Error(),
		)
	}

	global.UserSrvClient = proto.NewUserClient(userConn)
}
