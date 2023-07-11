package dao

import (
	"example.com/http_demo/model/dao/table"
	"example.com/http_demo/utils/zlog"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

//-------- 单表查询开始 -----------//

func GetAllOrders() (orders []*table.Order, err error) {
	orders = make([]*table.Order, 0)
	err = DB().Find(&orders).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func GetFirstOrderInTable() (order *table.Order, err error) {
	// 指针变量要先初始化, 直接传给Gorm的查询方法会空指针报错
	order = new(table.Order)
	//.First(&order) 传指针类型变量order的指针也行, 效果一样不用纠结
	err = DB().First(order).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func GetLastOrderInTable() (order *table.Order, err error) {
	order = new(table.Order)
	err = DB().Last(order).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func GetUserOrders(userId int64) (orders []*table.Order, err error) {
	orders = make([]*table.Order, 0)
	err = DB().Where("user_id = ?", userId).Find(&orders).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func GetUserPaidOrders(userId int64) (orders []*table.Order, err error) {
	orders = make([]*table.Order, 0)
	err = DB().Where("user_id = ?", userId).
		Where("state = ?", 2).
		// 这两个Where调用等价于 .Where("user_id = ? AND state = ?", userId, 2)
		// 等价于用Map查询
		// .Where(map[stirng]interface{} { // 不建议使用, 只能进行相等匹配, 且不能做参数类型限制
		//       user_id: userId,
		//       state: 2
		//  })
		Find(&orders).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func GetUserUnPaidOrders(userId int64) (orders []*table.Order, err error) {
	orders = make([]*table.Order, 0)
	err = DB().Where("user_id = ?", userId).
		Where("state IN (?)", []int{1, 3}).
		Find(&orders).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func GetUserNotPaidOrders(userId int64) (orders []*table.Order, err error) {
	orders = make([]*table.Order, 0)
	err = DB().Where("user_id = ?", userId).
		Not("state", 2).
		Find(&orders).Error
	// SELECT * FROM orders WHERE user_id = {userId} AND state != 2
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func GetUserNotFailedOrders(userId int64) (orders []*table.Order, err error) {
	orders = make([]*table.Order, 0)
	err = DB().Where("user_id = ?", userId).
		Not("state", []int{1, 3}).
		Find(&orders).Error
	// SELECT * FROM orders WHERE user_id = {userId} AND state NOT IN (1, 3)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

// OR 条件
func GetOrderUnPaidOrPayFailed() (orders []*table.Order, err error) {
	orders = make([]*table.Order, 0)
	err = DB().Where("state = ?", 1).
		Or("state = ?", 3).
		Find(&orders).Error
	// SELECT * FROM orders WHERE user_id = {userId} AND state NOT IN (1, 3)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func QueryOrderListForUser(userId int64, dataFrom, dataSize int, startDate, endDate string) (orders []*table.Order, orderTotal int, err error) {
	orders = make([]*table.Order, 0)

	db := DB().Where("user_id = ?", userId)
	// 根据参数值生成响应的SQL
	if startDate != "" {
		db = db.Where("created_at >= ?", startDate)
	}
	if endDate != "" {
		db = db.Where("created_at <= ?", endDate)
	}

	err = db.Debug().Order("id DESC").
		Offset(dataFrom).
		Limit(dataSize).
		Find(&orders).Count(&orderTotal).Error
	// 一般分页会要求返回当页数据和总条数, 这种写法会减少代码重复量, 就不需要再写下面的 QueryOrderCountForUser 函数了
	// SELECT * FROM orders WHERE ... ORDER BY id DESC LIMIT 10, OFFSET 0
	// SELECT COUNT(*) FROM orders WHERE ... ORDER BY id DESC LIMIT 10, OFFSET 0
	// (ORDER BY, LIMIT 对COUNT不起作用, 会正确返回WHERE条件对应的行数)
	zlog.Info("result debug", zap.Any("list", orders), zap.Any("total", orderTotal))
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	return
}

func QueryOrderCountForUser(userId int64, startDate, endDate string) (orderCount int, err error) {
	db := DB().Where("user_id = ?", userId)
	if startDate != "" {
		db = db.Where("created_at >= ?", startDate)
	}
	if endDate != "" {
		db = db.Where("created_at <= ?", startDate)
	}

	err = db.Count(&orderCount).Error
	return
}

type OrderSerialInfo struct {
	OrderNo string // 对应Scan扫描的结果集里的order_no 字段
	// 如果Scan扫描的结果字段与结构体字段不同, 需要使用tag进行标注
	OutOrderNo string `gorm:"column:platform_order_no"`
}

func GetOrderSerialInfo(orderNo string) (orderSerialInfo *OrderSerialInfo, err error) {

	orderSerialInfo = new(OrderSerialInfo)
	err = DB().Debug().Table(table.Order{}.TableName()).
		Select("order_no, platform_order_no").
		Where("order_no = ?", orderNo).
		Scan(&orderSerialInfo).Error
	// Scan 用于将查询结果集扫描进结构体或者结构体列表中去
	// 没查到数据会返回 gorm.ErrRecordNotFound 错误
	// Scan 可以用在SELECT 和 GROUP、JOIN 等指定返回字段的查询中, 用于获取结果--把查询结果集扫描到实参变量中存储
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	zlog.Info("GetOrderSerialInfo", zap.Any("orderSerialInfo", orderSerialInfo))

	return
}

func GetOrdersSerial(orderNo []string) (serialInfoList []*OrderSerialInfo, err error) {
	serialInfoList = make([]*OrderSerialInfo, 0)
	err = DB().Debug().Table(table.Order{}.TableName()).
		Select("order_no, platform_order_no").
		Where("order_no IN (?)", orderNo).
		Scan(&serialInfoList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return
}

type OrderStat struct {
	Date           string
	DayMoneyAmount int64 `gorm:"column:amount"`
}

func GetOrderStat() (stats []*OrderStat, err error) {
	stats = make([]*OrderStat, 0)
	err = DB().Debug().Table(table.Order{}.TableName()).
		Select("SUM(bill_money) AS amount, DATE(created_at) AS date").
		Group("DATE(created_at)").
		Having("amount > 0").Scan(&stats).Error // 把Group结果扫描进实现定义好的结构体列表
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return
}

func GetOrderStatMap() (statMap map[string]*OrderStat, err error) {
	// 使用Rows()返回结果集, 遍历结果集返回更定制化的结果
	rows, err := DB().Table(table.Order{}.TableName()).
		Select("SUM(bill_money) AS amount, DATE(created_at) AS date").
		Group("DATE(created_at)").
		Having("amount > 0").Rows()

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	var amount int64
	var date string
	statMap = map[string]*OrderStat{}
	for rows.Next() {
		err = rows.Scan(&amount, &date)

		if err != nil {
			return nil, err
		}

		statMap[date] = &OrderStat{
			Date:           date,
			DayMoneyAmount: amount,
		}
	}

	return statMap, err
}

//-------- 单表查询结束 -----------//
