package table

import "time"

// 订单表
type Order struct {
	BaseFields
	UserId          int64     `gorm:"column:user_id" json:"user_id"`                                      //用户ID
	BillMoney       int64     `gorm:"column:bill_money" json:"bill_money"`                                //订单金额（分）
	OrderNo         string    `gorm:"column:order_no;type:varchar(32)" json:"order_no"`                   //业务支付订单号
	PlatformOrderNo string    `gorm:"column:platform_order_no;type:varchar(32)" json:"platform_order_no"` //支付中心订单号
	OrderGoodsId    int64     `gorm:"column:order_goods_id" json:"order_goods_id"`                        //订单对应的商品ID
	State           int8      `gorm:"column:state;default:1" json:"state"`                                //1-待支付，2-支付成功，3-支付失败
	PaidAt          time.Time `gorm:"column:paid_at" json:"paid_at"`
}

func (Order) TableName() string {
	return "orders"
}
