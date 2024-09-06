package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"strings"
	"test-api/user-web/forms"
	"test-api/user-web/global"
	"time"
)

func GenerateSmsCode(witdh int) string {
	//动态生成验证码
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < witdh; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms(c *gin.Context) {
	sForm := forms.SendForm{}
	if err := c.ShouldBind(&sForm); err != nil {
		HandleValidatorError(c, err)
		return
	}

	//可以用阿里云，腾讯云等等都可以，这里就直接json返回了
	s := GenerateSmsCode(6)

	//保存验证码到redis
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.Redis.Host, global.ServerConfig.Redis.Port),
	})
	rdb.Set(context.Background(), sForm.Mobile, s, 300*time.Second)

	c.JSON(200, gin.H{
		"code": 200,
		"data": s,
	})
}
