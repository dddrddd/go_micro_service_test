package handler

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"test-service/goods_srv/global"
	"test-service/goods_srv/model"
	"test-service/goods_srv/proto"
)

func (s *GoodsServer) GetAllCategorysList(context.Context, *emptypb.Empty) (*proto.CategoryListResponse, error) {
	var categorys []model.Category
	//global.DB.Preload("SubCategory").Find(&categorys)//这样只能获取到以下一级的数据,并且获取了所有级别的数据
	global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	b, _ := json.Marshal(&categorys)
	return &proto.CategoryListResponse{JsonData: string(b)}, nil
}

func (s *GoodsServer) GetSubCategory(c context.Context, req *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	categoryListResponse := proto.SubCategoryListResponse{}
	var subCategoryResponse []*proto.CategoryInfoResponse
	var category model.Category
	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "商品分类不存在")
	}
	categoryListResponse.Info = &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		Level:          category.Level,
		IsTab:          category.IsTab,
		ParentCategory: category.ParentCategoryID,
	}

	var subCategorys []model.Category
	preloads := "SubCategory"
	if category.Level == 1 {
		preloads = "SubCategory.SubCategory"
	}
	global.DB.Where(&model.Category{Level: 1}).Preload(preloads).Find(&subCategorys)
	for _, subCategory := range subCategorys {
		subCategoryResponse = append(subCategoryResponse, &proto.CategoryInfoResponse{
			Id:             subCategory.ID,
			Name:           subCategory.Name,
			Level:          subCategory.Level,
			IsTab:          subCategory.IsTab,
			ParentCategory: subCategory.ParentCategoryID,
		})
	}
	categoryListResponse.SubCategory = subCategoryResponse
	return &categoryListResponse, nil
}
func (s *GoodsServer) CreateCategory(c context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	category := model.Category{}

	category.Name = req.Name
	category.Level = req.Level
	if category.Level != 1 {
		category.ParentCategoryID = req.ParentCategory
	}
	category.IsTab = req.IsTab
	global.DB.Create(&category)
	return &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		Level:          category.Level,
		IsTab:          category.IsTab,
		ParentCategory: category.ParentCategoryID,
	}, nil
}
func (s *GoodsServer) DeleteCategory(c context.Context, req *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Category{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	return &emptypb.Empty{}, nil
}
func (s *GoodsServer) UpdateCategory(c context.Context, req *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	var category model.Category
	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	if req.Name != "" {
		category.Name = req.Name
	}
	if req.ParentCategory != 0 {
		category.ParentCategoryID = req.ParentCategory
	}
	if req.Level != 0 {
		category.Level = req.Level
	}
	if req.IsTab {
		category.IsTab = req.IsTab
	}
	global.DB.Save(&category)
	return &emptypb.Empty{}, nil
}
