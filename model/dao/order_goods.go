package dao

import (
	"example.com/http_demo/model/dao/table"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

func GetAllOrderGoods() (orderGoodsList []*table.OrderGoods, err error) {
	orderGoodsList = make([]*table.OrderGoods, 0) // List 类型的参数必须先初始化， 直接.Find(&orderGoodsList)会空指针报错
	err = DB().Debug().Find(&orderGoodsList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

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
		//DB().Create(&orderGoods).Error //创建单个orderGoods
	}
	// 批量创建orderGoods
	// v1 版本不支持用Create 进行批量插入, v2 已支持
	// err = DB().Create(&orderGoodsList).Error
	BulkInsertOrderGoods(orderGoodsList)
	return err
}

// BulkInsertOrderGoods
// 目前 gorm  v1 并不支持批量插入这一功能， v2版本支持
// 在 v1.x 的情况下，如果硬要使用 gorm 做批量插入，需要使用db.Exec方法
// 可参见以下方式：
// https://stackoverflow.com/questions/12486436/how-do-i-batch-sql-statements-with-package-database-sql
func BulkInsertOrderGoods(unsavedRows []*table.OrderGoods) error {
	valueStrings := make([]string, 0, len(unsavedRows))
	valueArgs := make([]interface{}, 0, len(unsavedRows)*3)
	for _, row := range unsavedRows {
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, row.UserId)
		valueArgs = append(valueArgs, row.GoodsName)
		valueArgs = append(valueArgs, row.OrderId)
	}
	statement := fmt.Sprintf("INSERT INTO "+table.OrderGoods{}.TableName()+" (user_id, goods_name, order_id) VALUES %s",
		strings.Join(valueStrings, ","))

	err := DB().Exec(statement, valueArgs...).Error
	return err
}

func CreateOrderGoodsInTx(tx *gorm.DB, userId, orderId int64, goodsName string) error {
	orderGoods := &table.OrderGoods{
		UserId:    userId,
		GoodsName: goodsName,
		OrderId:   orderId,
	}
	err := tx.Create(orderGoods).Error
	return err
}
