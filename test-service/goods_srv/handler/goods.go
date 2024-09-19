package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"test-service/goods_srv/global"
	"test-service/goods_srv/model"
	"test-service/goods_srv/proto"
)

type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

func (s *GoodsServer) GoodsList(c context.Context, req *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
	goodsListResponse := &proto.GoodsListResponse{}
	var goods []model.Goods
	query := global.DB.Model(&model.Goods{})
	if req.KeyWords != "" {
		query = query.Where("name LIKE ?", "%"+req.KeyWords+"%")
	}
	if req.IsHot {
		query = query.Where("is_hot = true")
	}
	if req.IsNew {
		query = query.Where("is_new = true")
	}
	if req.PriceMin > 0 {
		query = query.Where("shop_price >= ?", req.PriceMin)
	}
	if req.PriceMax > 0 {
		query = query.Where("shop_price <= ?", req.PriceMax)
	}
	if req.Brand != 0 {
		query = query.Where("brand_id = ?", req.Brand)
	}

	if req.TopCategory != 0 {
		var category model.Category
		if result := global.DB.First(&category, req.TopCategory); result.Error != nil {
			return nil, status.Errorf(codes.NotFound, "商品分类不存在")
		}
		switch category.Level {
		case 1:
			query1 := global.DB.Model(&model.Category{}).Select("id").Where("parent_category_id = ?", req.TopCategory)
			query = query.Where("category_id IN (?)", query1)
		case 2:
			query1 := global.DB.Model(&model.Category{}).Select("id").Where("parent_category_id = ?", req.TopCategory)
			query = query.Where("category_id IN (?)", query1)
		case 3:
			query = query.Where("category_id = ?", req.TopCategory)
		}
	}

	var count int64
	query.Count(&count)
	goodsListResponse.Total = int32(count)
	result := query.Preload("Category").Preload("Brands").Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&goods)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, good := range goods {
		GoodsInfoResponse := &proto.GoodsInfoResponse{
			Id:              good.ID,
			CategoryId:      good.CategoryID,
			Name:            good.Name,
			GoodsSn:         good.GoodsSn,
			ClickNum:        good.ClickNum,
			SoldNum:         good.SoldNum,
			FavNum:          good.FavNum,
			MarketPrice:     good.MarketPrice,
			ShopPrice:       good.ShopPrice,
			GoodsBrief:      good.GoodsBrief,
			ShipFree:        good.ShipFree,
			Images:          good.Images,
			DescImages:      good.DescImages,
			GoodsFrontImage: good.GoodsFrontImage,
			IsNew:           good.IsNew,
			IsHot:           good.IsHot,
			OnSale:          good.OnSale,
			Category: &proto.CategoryBriefInfoResponse{
				Id:   good.Category.ID,
				Name: good.Category.Name,
			},
			Brand: &proto.BrandInfoResponse{
				Id:   good.Brands.ID,
				Name: good.Brands.Name,
				Logo: good.Brands.Logo,
			},
		}
		goodsListResponse.Data = append(goodsListResponse.Data, GoodsInfoResponse)
	}
	return goodsListResponse, nil
}

func (s *GoodsServer) BatchGetGoods(c context.Context, req *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
	goodsListResponse := &proto.GoodsListResponse{}
	var goods []model.Goods
	result := global.DB.Where(req.Id).Find(&goods)
	for _, good := range goods {
		GoodsInfoResponse := &proto.GoodsInfoResponse{
			Id:              good.ID,
			CategoryId:      good.CategoryID,
			Name:            good.Name,
			GoodsSn:         good.GoodsSn,
			ClickNum:        good.ClickNum,
			SoldNum:         good.SoldNum,
			FavNum:          good.FavNum,
			MarketPrice:     good.MarketPrice,
			ShopPrice:       good.ShopPrice,
			GoodsBrief:      good.GoodsBrief,
			ShipFree:        good.ShipFree,
			Images:          good.Images,
			DescImages:      good.DescImages,
			GoodsFrontImage: good.GoodsFrontImage,
			IsNew:           good.IsNew,
			IsHot:           good.IsHot,
			OnSale:          good.OnSale,
			Category: &proto.CategoryBriefInfoResponse{
				Id:   good.Category.ID,
				Name: good.Category.Name,
			},
			Brand: &proto.BrandInfoResponse{
				Id:   good.Brands.ID,
				Name: good.Brands.Name,
				Logo: good.Brands.Logo,
			},
		}
		goodsListResponse.Data = append(goodsListResponse.Data, GoodsInfoResponse)
	}
	goodsListResponse.Total = int32(result.RowsAffected)
	return goodsListResponse, nil
}

func (s *GoodsServer) GetGoodsDetail(c context.Context, req *proto.GoodInfoRequest) (*proto.GoodsInfoResponse, error) {
	var good model.Goods
	if result := global.DB.First(&good, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	goodsInfoResponse := &proto.GoodsInfoResponse{
		Id:              good.ID,
		CategoryId:      good.CategoryID,
		Name:            good.Name,
		GoodsSn:         good.GoodsSn,
		ClickNum:        good.ClickNum,
		SoldNum:         good.SoldNum,
		FavNum:          good.FavNum,
		MarketPrice:     good.MarketPrice,
		ShopPrice:       good.ShopPrice,
		GoodsBrief:      good.GoodsBrief,
		ShipFree:        good.ShipFree,
		Images:          good.Images,
		DescImages:      good.DescImages,
		GoodsFrontImage: good.GoodsFrontImage,
		IsNew:           good.IsNew,
		IsHot:           good.IsHot,
		OnSale:          good.OnSale,
		Category: &proto.CategoryBriefInfoResponse{
			Id:   good.Category.ID,
			Name: good.Category.Name,
		},
		Brand: &proto.BrandInfoResponse{
			Id:   good.Brands.ID,
			Name: good.Brands.Name,
			Logo: good.Brands.Logo,
		},
	}
	return goodsInfoResponse, nil
}

func (s *GoodsServer) CreateGoods(c context.Context, req *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}
	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}
	goods := model.Goods{
		Brands:          brand,
		BrandID:         brand.ID,
		Category:        category,
		CategoryID:      category.ID,
		Name:            req.Name,
		GoodsSn:         req.GoodsSn,
		MarketPrice:     req.MarketPrice,
		ShopPrice:       req.ShopPrice,
		GoodsBrief:      req.GoodsBrief,
		ShipFree:        req.ShipFree,
		Images:          req.Images,
		DescImages:      req.DescImages,
		GoodsFrontImage: req.GoodsFrontImage,
		IsHot:           req.IsHot,
		IsNew:           req.IsNew,
		OnSale:          req.OnSale,
	}
	global.DB.Create(&goods)
	return &proto.GoodsInfoResponse{
		Id: goods.ID,
	}, nil
}
func (s *GoodsServer) DeleteGoods(c context.Context, req *proto.DeleteGoodsInfo) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Brands{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	return &emptypb.Empty{}, nil
}
func (s *GoodsServer) UpdateGoods(c context.Context, req *proto.CreateGoodsInfo) (*emptypb.Empty, error) {
	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}
	var goods model.Goods
	if result := global.DB.First(&goods, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品不存在")
	}

	goods.Brands = brand
	goods.BrandID = brand.ID
	goods.Category = category
	goods.CategoryID = category.ID
	goods.Name = req.Name
	goods.GoodsSn = req.GoodsSn
	goods.MarketPrice = req.MarketPrice
	goods.ShopPrice = req.ShopPrice
	goods.ShipFree = req.ShipFree
	goods.GoodsBrief = req.GoodsBrief
	goods.Images = req.Images
	goods.DescImages = req.DescImages
	goods.GoodsFrontImage = req.GoodsFrontImage
	goods.IsHot = req.IsHot
	goods.OnSale = req.OnSale
	goods.IsNew = req.IsNew
	global.DB.Save(&goods)
	return &emptypb.Empty{}, nil
}
