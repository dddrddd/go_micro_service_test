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

func (s *GoodsServer) BannerList(c context.Context, req *emptypb.Empty) (*proto.BannerListResponse, error) {
	bannerListResponse := &proto.BannerListResponse{}
	var banners []model.Banner
	result := global.DB.Find(&banners)
	bannerListResponse.Total = int32(result.RowsAffected)

	var bannerResponses []*proto.BannerResponse
	for _, banner := range banners {
		bannerResponses = append(bannerResponses, &proto.BannerResponse{
			Id:    banner.ID,
			Image: banner.Image,
			Index: banner.Index,
			Url:   banner.Url,
		})
	}
	bannerListResponse.Data = bannerResponses
	return bannerListResponse, nil
}
func (s *GoodsServer) CreateBanner(c context.Context, req *proto.BannerRequest) (*proto.BannerResponse, error) {
	banner := &model.Banner{
		Image: req.Image,
		Index: req.Index,
		Url:   req.Url,
	}
	global.DB.Create(banner)
	return &proto.BannerResponse{Id: banner.ID}, nil
}
func (s *GoodsServer) DeleteBanner(c context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Banner{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "品牌不存在")
	}
	return &emptypb.Empty{}, nil
}
func (s *GoodsServer) UpdateBanner(c context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	banner := &model.Banner{}
	if result := global.DB.Delete(&model.Banner{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "品牌不存在")
	}
	if req.Image != "" {
		banner.Image = req.Image
	}

	if req.Url != "" {
		banner.Url = req.Url
	}

	if req.Index != 0 {
		banner.Index = req.Index
	}
	global.DB.Save(&banner)
	return &emptypb.Empty{}, nil
}
