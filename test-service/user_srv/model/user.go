package model

import (
	"time"
)

type BaseModel struct {
	ID        int32     `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	//DeletedAt gorm.DeletedAt //表示软删除
	IsDeleted bool //第二种方法
}

// User 密码需要加密，一般用md5算法
type User struct {
	BaseModel

	Password string     `gorm:"type:varchar(100);not null"`
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"` //index是索引
	Nickname string     `gorm:"type:varchar(20)"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;default:'male';type:varchar(6) comment 'female表示女，male表示男'"`
	Role     int        `gorm:"column:role;default:1;type:int comment '1表示普通用户，2表示管理员'"`
}
