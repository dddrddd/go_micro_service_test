package main

import (
	"context"
	"fmt"
	"test-service/goods_srv/handler"
	"test-service/goods_srv/proto"
)

func main() {
	ss := &proto.CreateGoodsInfo{
		Name:            "cjt",
		ShopPrice:       20.5,
		GoodsBrief:      "safsaf",
		ShipFree:        false,
		GoodsFrontImage: "telnet://gdoysjym.travel/nqmxsiinj",
		Stocks:          100,
		MarketPrice:     20.5,
		GoodsSn:         "JZW4WhF",
		CategoryId:      135476,
		BrandId:         624,
		Images:          []string{"https://www.baidu.com"},
		DescImages:      []string{"news://edlftux.ie/ecufrll"},
	}
	var t *handler.GoodsServer
	rsp, _ := t.CreateGoods(context.Background(), ss)
	fmt.Println(rsp)
}
