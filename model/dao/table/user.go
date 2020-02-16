package table

import "time"

type User struct {
	Id        int64     `gorm:"column:id;primary_key"`
	UserName  string    `gorm:"column:username"`
	Secret    string    `gorm:"column:secret;type:varchar(1000)"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName sets the insert table name for this struct type
func (m *User) TableName() string {
	return "users"
}