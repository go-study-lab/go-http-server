package logic

import (
	"errors"
	"example.com/http_demo/model/dao"
	"example.com/http_demo/model/dao/table"
	"example.com/http_demo/utils/zlog"
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
