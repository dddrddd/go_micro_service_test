package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"test-service/user_srv/config"
	"time"
)

var (
	DB            *gorm.DB
	ServiceConfig config.ServerConfig
	NacosConfig   config.NacosConfig
)

func init() {
	dsn := fmt.Sprintf("root:yujingpig@tcp(127.0.0.1:3306)/user_srv?charset=utf8mb4&parseTime=True&loc=Local")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{SlowThreshold: time.Second, LogLevel: logger.Info, Colorful: true},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

}
