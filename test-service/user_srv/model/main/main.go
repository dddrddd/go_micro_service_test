package main

import (
	"crypto/md5"
	"encoding/hex"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"test-service/user_srv/model"
	"time"
)

func genMD5(code string) string {
	md5Ctx := md5.New()
	_, _ = io.WriteString(md5Ctx, code)
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func main() {

	dsn := "root:yujingpig@tcp(127.0.0.1:3306)/user_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&model.User{})

	//var user model.User
	//options := &password.Options{16, 100, 50, sha512.New}
	//salt, encodedPwd := password.Encode("admin123", options)
	//newPassword := fmt.Sprintf("$pbkdf2-sha513$%s$%s", salt, encodedPwd)
	//for i := 0; i < 10; i++ {
	//	user = model.User{
	//		Nickname: fmt.Sprintf("bobby%d", i),
	//		Mobile:   fmt.Sprintf("1818181899%d", i),
	//		Password: newPassword,
	//	}
	//	global.DB.Save(&user)
	//}
}
