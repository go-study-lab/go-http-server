package logic

import (
	"errors"
	"example.com/http_demo/model/dao"
	"example.com/http_demo/model/dao/table"
	"example.com/http_demo/utils/zlog"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func InitOrderGoodsData() error {
	if err := CheckOrdersInitialization(); err != nil {
		return err
	}
	err := dao.InitOrderGoodsFromOrder()
	if err != nil {
		zlog.Error("order_init_error", zap.Error(err))
	}
	return err
}

func CheckOrdersInitialization() error {
	orderGoodsList, err := dao.GetAllOrderGoods()
	if err != nil {
		return err
	}

	if len(orderGoodsList) > 0 {
		err = errors.New("OrderGoods 已初始化, 请勿重复操作")
		return err
	}
	return nil
}

func QueryUserOrdersInPage(userId int64, page int, startDate, endDate string) (orders []*table.Order, orderTotal int, err error) {
	dataSize := 10
	dataFrom := (page - 1) * dataSize
	orders, orderTotal, err = dao.QueryOrderListForUser(userId, dataFrom, dataSize, startDate, endDate)
	if err != nil {
		zlog.Error("QueryUserOrdersError", zap.Error(err))
		return nil, 0, err
	}
	return
}

func SetOrderSuccessAndCreateGoods(userId int64, orderNo string) (err error) {

	// 开启事务并自动提交
	err = dao.DB().Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		order, err := dao.GetOrderByNoInTx(tx, orderNo)
		if err != nil || order.UserId != userId {
			zlog.Error("error_detail", zap.Error(err), zap.Any("userId", userId), zap.Any("orderNo", orderNo))
			// 返回任何错误都会回滚事务
			return errors.New("未找到对应的订单")
		}

		err = dao.SetOrderStatePaySuccessInTx(tx, orderNo)
		if err != nil {
			return err
		}

		err = dao.CreateOrderGoodsInTx(tx, userId, order.Id, "商品名InTx")

		return err
	})

	return
}

func SetOrderSuccessAndCreateGoodsInHandTx(userId int64, orderNo string) (err error) {
	// 手动开启事务
	tx := dao.DB().Begin()
	panicked := true
	defer func() { // 控制事务的提交和回滚, 保证事务的完整性
		// db.Transaction 内部其实就是这么实现的
		if err != nil || panicked {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	order, err := dao.GetOrderByNoInTx(tx, orderNo)
	if err != nil || order.UserId != userId {
		zlog.Error("error_detail", zap.Error(err), zap.Any("userId", userId), zap.Any("orderNo", orderNo))
		// 返回任何错误都会回滚事务
		return errors.New("未找到对应的订单")
	}

	err = dao.SetOrderStatePaySuccessInTx(tx, orderNo)
	if err != nil {
		return err
	}

	err = dao.CreateOrderGoodsInTx(tx, userId, order.Id, "商品名InTx")
	if err != nil {
		return err
	}
	// 没有错误把panicked设置为false,  代表着程序正常执行完毕, 事务提交
	panicked = false
	return err
}
