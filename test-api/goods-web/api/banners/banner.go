package banners

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"strconv"
	"test-api/goods-web/api"
	"test-api/goods-web/forms"
	"test-api/goods-web/global"
	"test-api/goods-web/proto"
)

func BannerList(c *gin.Context) {
	rsp, err := global.GoodsSrvClient.BannerList(context.Background(), &emptypb.Empty{})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["index"] = value.Index
		reMap["image"] = value.Image
		reMap["url"] = value.Url
		result = append(result, reMap)
	}
	c.JSON(http.StatusOK, result)
}

func NewBanner(c *gin.Context) {
	bannerForm := forms.BannerForm{}
	if err := c.ShouldBind(&bannerForm); err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	rsp, err := global.GoodsSrvClient.CreateBanner(context.Background(), &proto.BannerRequest{
		Index: int32(bannerForm.Index),
		Image: bannerForm.Image,
		Url:   bannerForm.Url,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	result := make(map[string]interface{})
	result["id"] = rsp.Id
	result["index"] = rsp.Index
	result["image"] = rsp.Image
	result["url"] = rsp.Url
	c.JSON(http.StatusOK, result)
}

func UpdateBanner(c *gin.Context) {
	bannerForm := forms.BannerForm{}
	if err := c.ShouldBind(&bannerForm); err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.UpdateBanner(context.Background(), &proto.BannerRequest{
		Id:    int32(i),
		Index: int32(bannerForm.Index),
		Url:   bannerForm.Url,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

func DeleteBanner(c *gin.Context) {
	id := c.Query("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.DeleteBanner(context.Background(), &proto.BannerRequest{
		Id: int32(i),
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}
