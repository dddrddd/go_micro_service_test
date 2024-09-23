package forms

type BannerForm struct {
	Image string `json:"image" form:"image" binding:"url"`
	Index int    `json:"index" form:"index" binding:"required"`
	Url   string `json:"url" form:"url" binding:"url"`
}
