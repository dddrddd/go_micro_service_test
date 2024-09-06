package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func genMD5(code string) string {
	md5Ctx := md5.New()
	_, _ = io.WriteString(md5Ctx, code)
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func main() {
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
