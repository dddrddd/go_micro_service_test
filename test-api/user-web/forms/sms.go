package forms

type SendForm struct {
	Mobile string `json:"mobile" binding:"required,mobile"`
	Type   uint   `json:"type" binding:"required,oneof=1 2"` //表示发送短信的类型,1表示注册，2找回密码
}
