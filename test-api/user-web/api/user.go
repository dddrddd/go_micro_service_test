package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"strings"
	"test-api/user-web/forms"
	"test-api/user-web/global"
	"test-api/user-web/global/response"
	"test-api/user-web/middlewares"
	"test-api/user-web/models"
	"test-api/user-web/proto"
	"time"
)

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for filed, err := range fileds {
		rsp[filed[strings.Index(filed, ".")+1:]] = err
	}
	return rsp
}

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误",
				})
			}
			return
		}
	}
	// 如果err不是一个gRPC的错误
	c.JSON(http.StatusInternalServerError, gin.H{
		"msg": "未知错误",
	})
	return
}
func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}

func GetUserList(c *gin.Context) {

	claims, _ := c.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("id:%d", currentUser.ID)
	pn := c.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := c.DefaultQuery("pSize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList]查询用户列表失败", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		user := response.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			BirthDay: response.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		result = append(result, user)
	}
	c.JSON(http.StatusOK, result)
}

func PassWordLogin(c *gin.Context) {
	//表单验证
	passwordLoginForm := forms.PassWordLoginForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	//这里可以加一个图形验证
	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, true) {
		c.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
		return
	}

	//登录逻辑
	rsp, err := global.UserSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"mobile": "用户不存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"mobile": "登陆失败",
				})
			}
			return
		}
	} else {
		//检查密码
		if passRsp, passErr := global.UserSrvClient.CheckPassword(context.Background(), &proto.CheckInfo{
			Password:          passwordLoginForm.Password,
			EncryptedPassword: rsp.Password,
		}); passErr != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"password": "登陆失败",
			})
		} else {
			if passRsp.Success {
				//生成token
				j := middlewares.NewJWT()
				claims := models.CustomClaims{
					ID:          uint(rsp.Id),
					Nickname:    rsp.NickName,
					AuthorityId: uint(rsp.Role),
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),              //签名的生效时间
						ExpiresAt: time.Now().Unix() + 60*60*24*7, //30天过期
						Issuer:    "go-learn",
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"msg": "生成token失败",
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"id":        rsp.Id,
					"nickName":  rsp.NickName,
					"token":     token,
					"expiresAt": (time.Now().Unix() + 60*60*24*7) * 1000,
				})
			} else {
				c.JSON(http.StatusBadRequest, map[string]interface{}{
					"msg": "登陆失败",
				})
			}

		}
	}

}

func Register(c *gin.Context) {
	registerForm := forms.RegisterForm{}
	if err := c.ShouldBind(&registerForm); err != nil {
		HandleValidatorError(c, err)
		return
	}

	//验证码校验
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.Redis.Host, global.ServerConfig.Redis.Port),
	})
	value, err := rdb.Get(context.Background(), registerForm.Mobile).Result()
	if errors.Is(err, redis.Nil) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "验证码错误",
		})
		return
	} else {
		if value != registerForm.Code {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "验证码错误",
			})
			return
		}
	}

	user, err := global.UserSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		Mobile:   registerForm.Mobile,
		NickName: registerForm.Mobile,
		Password: registerForm.Password,
	})
	if err != nil {
		zap.S().Errorw("[Register]新建用户服务失败", err.Error())
		HandleGrpcErrorToHttp(err, c)
		return
	}
	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:          uint(user.Id),
		Nickname:    user.NickName,
		AuthorityId: uint(user.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),              //签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, //30天过期
			Issuer:    "go-learn",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        user.Id,
		"nickName":  user.NickName,
		"token":     token,
		"expiresAt": (time.Now().Unix() + 60*60*24*7) * 1000,
	})
}
