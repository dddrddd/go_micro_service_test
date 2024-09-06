package api

//图片验证码
import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, answer, err := cp.Generate()
	if err != nil {
		zap.S().Errorw("生成验证码错误", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成验证码错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath":   b64s,
		"answer":    answer,
	})
	//store.Verify(id, answer, true)
}
