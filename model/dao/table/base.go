package table

import "time"

type BaseFields struct {
	Id        int64     `gorm:"column:id;primary_key" json:"id"`     //自增ID
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` //创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` //更新时间
}
