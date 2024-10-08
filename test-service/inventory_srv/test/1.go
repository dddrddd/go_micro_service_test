package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"test-service/inventory_srv/proto"
)

var invClient proto.InventoryClient
var conn *grpc.ClientConn

func Init() {
	conn, _ = grpc.Dial("localhost:58431", grpc.WithInsecure())
	invClient = proto.NewInventoryClient(conn)

}

func SetInv(goodsId, Num int32) {
	_, err := invClient.SetInv(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
		Number:  Num,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("111")
}

func INV(goodsId int32) {
	rsp, err := invClient.InvDetail(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func Sell() {
	_, err := invClient.Sell(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{
				GoodsId: 421,
				Number:  1,
			},
			{
				GoodsId: 422,
				Number:  2,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("211")
}

func reback() {
	_, err := invClient.Reback(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{
				GoodsId: 421,
				Number:  1,
			},
			{
				GoodsId: 422,
				Number:  2,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("211")
}
func main() {
	Init()
	for i := 421; i <= 840; i++ {
		SetInv(int32(i), 100)
	}

	//INV(421)
	//Sell()
	//reback()
	err := conn.Close()
	if err != nil {
		return
	}
}
