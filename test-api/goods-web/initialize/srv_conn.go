package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"test-api/goods-web/global"
	"test-api/goods-web/proto"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%s/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsInfo.Name), //这一部分使用了"github.com/mbobakov/grpc-consul-resolver"，固定格式，标签可以参考官方文档
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), //这部分也是
	)
	if err != nil {
		zap.S().Fatalf("用户服务失败")
	}
	global.GoodsSrvClient = proto.NewGoodsClient(userConn)
}
