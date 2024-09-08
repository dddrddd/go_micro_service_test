package model

type Category struct {
	BaseModel
	Name             string `gorm:"type:varchar(255);not null"`
	Level            int32  `gorm:"type:int;not null;default:1;comment:'共三级'"`
	IsTab            bool   `gorm:"not null;default:false"`
	ParentCategoryID int32
	ParentCategory   *Category
}

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(255);not null"`
	Logo string `gorm:"type:varchar(255);default:'';not null"`
}

type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32    `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category `gorm:"foreignKey:CategoryID;references:ID"`
	BrandID    int32    `gorm:"type:int;index:idx_category_brand,unique"`
	Brands     Brands   `gorm:"foreignKey:BrandID;references:ID"`
}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(255);not null"`
	Url   string `gorm:"type:varchar(255);not null"`
	Index int32  `gorm:"type:int;not null;default:1"`
}

type Goods struct {
	BaseModel
	CategoryID int32    `gorm:"type:int;not null"`
	Category   Category `gorm:"foreignKey:CategoryID;references:ID"`
	BrandID    int32    `gorm:"type:int;not null"`
	Brands     Brands   `gorm:"foreignKey:BrandID;references:ID"`

	OnSale   bool `gorm:"not null;default:false"`
	ShipFree bool `gorm:"not null;default:false"`
	IsNew    bool `gorm:"not null;default:false"`
	IsHot    bool `gorm:"not null;default:false"`

	Name            string   `gorm:"type:varchar(255);not null"`
	GoodsSn         string   `gorm:"type:varchar(255);not null;commit:'商品编号'"`
	ClickNum        int32    `gorm:"type:int;not null;default:0"`
	SoldNum         int32    `gorm:"type:int;not null;default:0"`
	FavNum          int32    `gorm:"type:int;not null;default:0"`
	MarketPrice     float32  `gorm:"not null"`
	ShopPrice       float32  `gorm:"not null"`
	GoodsBrief      string   `gorm:"type:varchar(255);not null"`
	Images          GormList `gorm:"type:varchar(255);not null"`
	DescImages      GormList `gorm:"type:varchar(255);not null"`
	GoodsFrontImage string   `gorm:"type:varchar(255);not null"`
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}
