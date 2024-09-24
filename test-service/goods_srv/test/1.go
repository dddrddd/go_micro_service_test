package main

import (
	"context"
	"fmt"
	"test-service/goods_srv/handler"
	"test-service/goods_srv/proto"
)

func main() {
	ss := &proto.CategoryBrandRequest{
		Id:         25799,
		CategoryId: 130366,
		BrandId:    614,
	}
	var t *handler.GoodsServer
	rsp, err := t.UpdateCategoryBrand(context.Background(), ss)
	fmt.Println(rsp, err)
}
