package brands

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test-api/goods-web/api"
	"test-api/goods-web/forms"
	"test-api/goods-web/global"
	"test-api/goods-web/proto"
)

func BrandList(c *gin.Context) {
	pn := c.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := c.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := global.GoodsSrvClient.BrandList(context.Background(), &proto.BrandFilterRequest{
		Pages:       int32(pnInt),
		PagePerNums: int32(pSizeInt),
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	result := make([]interface{}, 0)

	for _, value := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["name"] = value.Name
		reMap["logo"] = value.Logo
		result = append(result, reMap)
	}
	finalMap := make(map[string]interface{})
	finalMap["total"] = rsp.Total
	finalMap["data"] = result

	c.JSON(http.StatusOK, finalMap)
}

func NewBrand(c *gin.Context) {
	brandForm := forms.BrandForm{}
	if err := c.ShouldBind(&brandForm); err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	rsp, err := global.GoodsSrvClient.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: brandForm.Name,
		Logo: brandForm.Logo,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	request := make(map[string]interface{})
	request["id"] = rsp.Id
	request["name"] = rsp.Name
	request["logo"] = rsp.Logo
	c.JSON(http.StatusOK, request)
}

func UpdateBrand(c *gin.Context) {
	brandForm := forms.BrandForm{}
	if err := c.ShouldBind(&brandForm); err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.UpdateBrand(context.Background(), &proto.BrandRequest{
		Id:   int32(i),
		Name: brandForm.Name,
		Logo: brandForm.Logo,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})

}

func DeleteBrand(c *gin.Context) {
	id := c.Query("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.DeleteBrand(context.Background(), &proto.BrandRequest{
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

func GetCategoryBrandList(c *gin.Context) {
	id := c.Query("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	rsp, err := global.GoodsSrvClient.GetCateGoryBrandList(context.Background(), &proto.CategoryInfoRequest{
		Id: int32(i),
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["name"] = value.Name
		reMap["logo"] = value.Logo
		result = append(result, reMap)
	}
	c.JSON(http.StatusOK, result)
}

func CategoryBrandList(c *gin.Context) {
	pn := c.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := c.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)
	rep, err := global.GoodsSrvClient.CategoryBrandList(context.Background(), &proto.CategoryBrandFilterRequest{
		Pages:       int32(pnInt),
		PagePerNums: int32(pSizeInt),
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	reMap := map[string]interface{}{}
	result := make([]interface{}, 0)

	for _, value := range rep.Data {
		reMap = make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["category"] = map[string]interface{}{
			"name": value.Category.Name,
			"id":   value.Category.Id,
		}
		reMap["brand"] = map[string]interface{}{
			"name": value.Brand.Name,
			"id":   value.Brand.Id,
			"logo": value.Brand.Logo,
		}
		result = append(result, reMap)
	}

	finalMap := make(map[string]interface{})
	finalMap["total"] = rep.Total
	finalMap["data"] = result

	c.JSON(http.StatusOK, finalMap)
}

func NewCategoryBrand(c *gin.Context) {
	categoryBrandForm := forms.CategoryBrandForm{}
	if err := c.ShouldBind(&categoryBrandForm); err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	rsp, err := global.GoodsSrvClient.CreateCategoryBrand(context.Background(), &proto.CategoryBrandRequest{
		CategoryId: int32(categoryBrandForm.CategoryId),
		BrandId:    int32(categoryBrandForm.BrandId),
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	response := make(map[string]interface{})
	response["id"] = rsp.Id
	c.JSON(http.StatusOK, response)
}

func UpdateCategoryBrand(c *gin.Context) {
	categoryBrandForm := forms.CategoryBrandForm{}
	if err := c.ShouldBind(&categoryBrandForm); err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.UpdateCategoryBrand(context.Background(), &proto.CategoryBrandRequest{
		Id:         int32(i),
		CategoryId: int32(categoryBrandForm.CategoryId),
		BrandId:    int32(categoryBrandForm.BrandId),
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

func DeleteCategoryBrand(c *gin.Context) {
	id := c.Query("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.DeleteCategoryBrand(context.Background(), &proto.CategoryBrandRequest{Id: int32(i)})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}
