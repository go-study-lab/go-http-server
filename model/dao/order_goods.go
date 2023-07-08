package dao

import (
	"example.com/http_demo/model/dao/table"
	"strconv"
)

func InitOrderGoodsFromOrder() error {
	orders, err := GetAllOrders()
	if err != nil {
		return err
	}
	orderGoodsList := make([]*table.OrderGoods, 0)
	for index, order := range orders {
		orderGoods := &table.OrderGoods{
			UserId:    order.UserId,
			GoodsName: "商品" + strconv.Itoa(index+1),
			OrderId:   order.Id,
		}
		orderGoodsList = append(orderGoodsList, orderGoods)
	}
	// 批量创建orderGoods
	err = DB().Create(&orderGoodsList).Error
	return err
}
