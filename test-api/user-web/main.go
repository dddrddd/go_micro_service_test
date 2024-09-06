package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"test-api/user-web/global"
	"test-api/user-web/initialize"
	"test-api/user-web/utils"

	myvalidator "test-api/user-web/validator"
)

func main() {
	initialize.InitLogger()

	initialize.InitConfig()

	Router := initialize.Routers()

	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	initialize.InitSrvConn()

	viper.AutomaticEnv()
	//如果是本地开发环境，端口号固定，线上环境自动获取端口号
	debug := viper.GetBool("test_debug")
	if !debug {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}

	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())

			return t
		})
	}

	port := global.ServerConfig.Port
	//logger,_ := zap.NewProduction()
	//defer logger.Sync()
	//sugar := logger.Sugar()

	zap.S().Debugf("启动服务器，端口：%d", port)
	//S()可以省略代码，直接获得全局的suger,可以让我们设置一个全局的logger在里面，这样就能获得sugerLogger的所有功能了，也不用过多配置了
	//也可以用l，就相当于得到logger，s和l可以省略我们进行互斥操作
	if err := Router.Run(fmt.Sprintf(":%v", port)); err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}
}
