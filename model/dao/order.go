package dao

import (
	"example.com/http_demo/model/dao/table"
	"github.com/jinzhu/gorm"
)

func GetAllOrders() (orders []*table.Order, err error) {
	orders = make([]*table.Order, 0) // List 类型的参数必须先初始化， 直接.Find(&orders)会空指针报错
	err = DB().Find(&orders).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}
