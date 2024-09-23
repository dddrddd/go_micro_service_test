package category

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"strconv"
	"test-api/goods-web/api"
	"test-api/goods-web/forms"
	"test-api/goods-web/global"
	"test-api/goods-web/proto"
)

func List(c *gin.Context) {
	r, err := global.GoodsSrvClient.GetAllCategorysList(context.Background(), &emptypb.Empty{})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}

	var data []map[string]interface{}
	err = json.Unmarshal([]byte(r.JsonData), &data)
	if err != nil {
		zap.S().Errorw("[List] 查询分类列表失败:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "分类数据解析失败"})
		return
	}

	// 正常返回 JSON 数据
	c.JSON(http.StatusOK, data)
}

func Detail(c *gin.Context) {
	id := c.Query("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	reMap := make(map[string]interface{})
	subCategorys := make([]interface{}, 0)
	if r, err := global.GoodsSrvClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: int32(i),
	}); err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	} else {
		for _, value := range r.SubCategory {
			subCategorys = append(subCategorys, map[string]interface{}{
				"id":              value.Id,
				"name":            value.Name,
				"level":           value.Level,
				"parent_category": value.ParentCategory,
				"is_tab":          value.IsTab,
			})
		}
		reMap["id"] = r.Info.Id
		reMap["name"] = r.Info.Name
		reMap["level"] = r.Info.Level
		reMap["parent_category"] = r.Info.ParentCategory
		reMap["is_tab"] = r.Info.IsTab

		c.JSON(http.StatusOK, reMap)
	}
}

func New(c *gin.Context) {
	categoryForm := forms.CategoryForm{}
	if err := c.ShouldBind(&categoryForm); err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}

	rsp, err := global.GoodsSrvClient.CreateCategory(context.Background(), &proto.CategoryInfoRequest{
		Name:           categoryForm.Name,
		Level:          categoryForm.Level,
		ParentCategory: categoryForm.ParentCategory,
		IsTab:          *categoryForm.IsTab,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	request := make(map[string]interface{})
	request["id"] = rsp.Id
	request["name"] = rsp.Name
	request["level"] = rsp.Level
	request["parent_category"] = rsp.ParentCategory
	request["is_tab"] = rsp.IsTab

	c.JSON(http.StatusOK, request)
}

func Delete(c *gin.Context) {
	id := c.Query("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.DeleteCategory(context.Background(), &proto.DeleteCategoryRequest{
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

func Update(c *gin.Context) {
	categoryForm := forms.UpdateCategoryForm{}
	if err := c.ShouldBindJSON(&categoryForm); err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}

	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	request := &proto.CategoryInfoRequest{
		Id:   int32(i),
		Name: categoryForm.Name,
	}
	if categoryForm.IsTab != nil {
		request.IsTab = *categoryForm.IsTab
	}
	_, err = global.GoodsSrvClient.UpdateCategory(context.Background(), request)
	if err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}
