package handler

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"sync"
	"test-service/inventory_srv/global"
	"test-service/inventory_srv/model"
	"test-service/inventory_srv/proto"
)

type InventoryServer struct {
	proto.UnimplementedInventoryServer
}

func (*InventoryServer) SetInv(c context.Context, req *proto.GoodsInvInfo) (*emptypb.Empty, error) {
	var inv model.Inventory
	global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv)
	if inv.Goods == 0 {
		inv.Goods = req.GoodsId
	}
	inv.Stocks = req.Number
	global.DB.Save(&inv)
	return &emptypb.Empty{}, nil
}

func (*InventoryServer) InvDetail(c context.Context, req *proto.GoodsInvInfo) (*proto.GoodsInvInfo, error) {
	var inv model.Inventory
	if result := global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "库存信息不存在")
	}
	return &proto.GoodsInvInfo{
		GoodsId: inv.Goods,
		Number:  inv.Stocks,
	}, nil

}

var m sync.Mutex

func (*InventoryServer) Sell(c context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	tx := global.DB.Begin()
	//m.Lock()
	for _, good := range req.GoodsInfo {
		var inv model.Inventory
		//if result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where(&model.Inventory{Goods: good.GoodsId}).First(&inv); result.RowsAffected == 0 {
		//	tx.Rollback()
		//	return nil, status.Errorf(codes.InvalidArgument, "库存信息不存在")
		//}
		for {
			if result := global.DB.Where(&model.Inventory{Goods: good.GoodsId}).First(&inv); result.RowsAffected == 0 {
				tx.Rollback()
				return nil, status.Errorf(codes.InvalidArgument, "库存信息不存在")
			}
			if inv.Stocks < good.Number {
				tx.Rollback()
				return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
			}
			inv.Stocks -= good.Number
			if result := tx.Model(&model.Inventory{}).Select("Stocks", "Version").Where("goods = ? and version = ?", good.GoodsId, inv.Version).Updates(model.Inventory{Stocks: inv.Stocks, Version: inv.Version + 1}); result.RowsAffected == 0 {
				zap.S().Info("库存扣减失败")
			} else {
				break
			}
		}
		//tx.Save(&inv)
	}
	tx.Commit()
	//m.Unlock()
	return &emptypb.Empty{}, nil
}

func (*InventoryServer) Reback(c context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	tx := global.DB.Begin()
	//todo 处理并发场景
	for _, good := range req.GoodsInfo {
		var inv model.Inventory
		if result := global.DB.Where(&model.Inventory{Goods: good.GoodsId}).First(&inv); result.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Errorf(codes.InvalidArgument, "库存信息不存在")
		}
		inv.Stocks += good.Number
		tx.Save(&inv)
	}
	tx.Commit()
	return &emptypb.Empty{}, nil
}
