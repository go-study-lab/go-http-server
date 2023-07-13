package dao

import (
	"example.com/http_demo/model/dao/table"
	"example.com/http_demo/utils/zlog"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"time"
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
		Where("state = ?", table.ORDER_STATE_PAY_SUCCESS).
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
		Where("state IN (?)", []int{table.ORDER_STATE_UNPAID, table.ORDER_STATE_PAY_FAILD}).
		Find(&orders).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func GetUserNotPaidOrders(userId int64) (orders []*table.Order, err error) {
	orders = make([]*table.Order, 0)
	err = DB().Where("user_id = ?", userId).
		Not("state", table.ORDER_STATE_PAY_SUCCESS).
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
		Not("state", []int{table.ORDER_STATE_UNPAID, table.ORDER_STATE_PAY_FAILD}).
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
	err = DB().Where("state = ?", table.ORDER_STATE_UNPAID).
		Or("state = ?", table.ORDER_STATE_PAY_FAILD).
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
	db := DB().Model(table.Order{}).
		Where("user_id = ?", userId)
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
	// Debug() 用于调试时打印出执行的SQL
	// Scan操作时使用.Model() 会告诉GORM要在哪个Model对象上执行操作, GORM会自动找到对应的表名
	err = DB().Debug().Model(table.Order{}).
		// DB().Model(table.Order{}) 也可以写成 DB().Table(table.Order{}.TableName())
		//.Table()是更底层的写法，直接指定SQL的表名, 一般在关联查询里使用这种写法
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

//db.Model() 和 db.Table()的区别
//https://stackoverflow.com/questions/67550966/go-gorm-difference-between-db-model-and-db-table-query

func GetOrdersSerial(orderNo []string) (serialInfoList []*OrderSerialInfo, err error) {
	serialInfoList = make([]*OrderSerialInfo, 0)
	err = DB().Debug().Model(table.Order{}).
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
	err = DB().Debug().Model(table.Order{}).
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
	defer rows.Close()
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

//-------- Update操作 -----------//

// GORM 的时间跟踪
// GORM 会对模型中的三个时间字段进行跟踪, 分别是CreatedAt, UpdatedAt, DeletedAt
// 如果模型有CreatedAt, UpdatedAt 创建记录时都会被赋予当时的时间, 更新记录时GORM会总动更新UpdatedAt的时间
// 如果模型有DeletedAt字段, 调用Delete删除该记录时，将会设置DeletedAt字段为当前时间，而不是直接将记录从数据库中删除。

func SetOrderStatePaySuccess(orderNo string) error {
	updates := map[string]interface{}{
		"state":   table.ORDER_STATE_PAY_SUCCESS,
		"paid_at": time.Now(),
	}
	// 执行SQL UPDATE orders SET state = 2, paid_at = '当前时间(yyyy-mm-dd hh:ii:ss)', updated_at = '当前时间' WHERE order_no = {orderNo}
	// 使用 map 更新多个属性，只会更新其中有变化的属性
	err := DB().Model(table.Order{}).
		Where("order_no = ?", orderNo).
		Update(updates).Error
	// Update方法里也可以使用结构体 Update(&table.Order{State: 2, PaidAt: time.Now()})
	// 使用 struct 更新多个属性，只会更新其中有变化且为非零值的字段
	// 另外可以通过Update(updates).RowsAffected 获取更新的记录数
	return err
}

//-------- Update操作结束 -----------//

//-------- Join 查询 -----------//

type UserGoods struct {
	UserId     int64
	OrderNo    string
	GoodsId    int64
	GoodsName  string
	OrderState int `gorm:"column:state"`
}

func GetUserOrderGoods(userId int64) (goodsList []*UserGoods, err error) {
	tableOrder := table.Order{}.TableName()
	tableGoods := table.OrderGoods{}.TableName()

	goodsList = make([]*UserGoods, 0)

	err = DB().Debug().Table(tableOrder+" AS o").
		Joins("INNER JOIN "+tableGoods+" AS og ON o.id = og.order_id").
		Select("o.user_id, o.order_no, o.state, og.id as goods_id, og.goods_name").
		Where("o.user_id = ?", userId).Scan(&goodsList).Error
	// Scan, Find方法的参数是Slice类型时, 必须传递Slice的指针, 否则在其中填充数据后因为底层数组变更,导致内外部Slice指向的底层数组不一致
	zlog.Info("debug data", zap.Any("data", goodsList))

	return
}

// GetUserOrderGoodsMap 返回以GoodsId为Key, GoodsName为值的Map
// 主要是为了演示对JOIN查询结果集的逐行Scan操作
func GetUserOrderGoodsMap(userId int64) (userOrderGoodsMap map[int64]string, err error) {
	tableOrder := table.Order{}.TableName()
	tableGoods := table.OrderGoods{}.TableName()

	rows, err := DB().Table(tableOrder+" AS o").
		Joins("INNER JOIN "+tableGoods+" AS og ON o.id = og.order_id").
		Select("og.id as goods_id, og.goods_name").
		Where("o.user_id = ?", userId).Rows()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	defer rows.Close()
	userOrderGoodsMap = make(map[int64]string)
	var (
		GoodsId   int64
		GoodsName string
	)
	for rows.Next() {
		//err := rows.Scan(&result)
		// 遍历结果集的记录时Scan只能往多个基础类型变量里扫描列对应的数据
		// 有几列就对应几个变量, 不能往结构体里扫描, 否则会报错
		// 比如本例中如果把连表查询的结果逐行Scan到一个结构体变量, 会有如下错误
		// sql: expected 2 destination arguments in Scan, not 1
		// 综合起来更推荐上一种把整个结果集Scan进结构体列表的方式拿到连表查询的结果集.
		err := rows.Scan(&GoodsId, &GoodsName)
		if err != nil {
			return nil, err
		}

		userOrderGoodsMap[GoodsId] = GoodsName
	}

	return
}
