package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type GormList []string //自定义数据类型在gorm中使用
func (l GormList) Value() (driver.Value, error) {
	return json.Marshal(l)
}
func (l *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &l)
}

type BaseModel struct {
	ID        int32     `gorm:"primary_key"` //使用int32以避免作为外键使用时的冲突
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	//DeletedAt gorm.DeletedAt //表示软删除
	IsDeleted bool `gorm:"column:is_deleted"` //第二种方法
}
