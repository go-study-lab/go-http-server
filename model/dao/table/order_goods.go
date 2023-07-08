package table

// 订单商品表
type OrderGoods struct {
	BaseFields
	UserId    int64  `gorm:"column:user_id" json:"user_id"`       //用户ID
	GoodsName string `gorm:"column:goods_name" json:"goods_name"` //商品名
	OrderId   int64  `gorm:"column:order_id" json:"order_id"`     //商品的订单ID
}

func (OrderGoods) TableName() string {
	return "order_goods"
}
