package forms

type PassWordLoginForm struct {
	Mobile    string `json:"mobile" binding:"required,mobile"` //手机号码格式需要自定义，这个mobile对应验证里面的那个参数
	Password  string `json:"password" binding:"required,min=5,max=20"`
	Captcha   string `json:"captcha" binding:"required" form:"captcha"`
	CaptchaId string `json:"captcha_id" binding:"required" form:"captcha_id"`
}
type RegisterForm struct {
	Mobile   string `json:"mobile" binding:"required,mobile"`
	Password string `json:"password" binding:"required,min=5,max=20"`
	Code     string `json:"code" binding:"required,min=6,max=6" form:"code"`
}
