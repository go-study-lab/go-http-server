package order

import (
	"example.com/http_demo/logic"
	"example.com/http_demo/model/dao"
	"fmt"
	"net/http"
)

func GoodsInit(w http.ResponseWriter, r *http.Request) {

	err := logic.InitOrderGoodsData()
	if err != nil {
		fmt.Fprintln(w, "初始化OrderGoods失败!", "原因:"+err.Error())
	} else {
		fmt.Fprintln(w, "初始化OrderGoods成功!")
	}
	return
}

func SingeTableQuery(w http.ResponseWriter, r *http.Request) {
	// 根据主键查询第一条记录
	data, _ := dao.GetFirstOrderInTable()
	fmt.Fprintln(w, "Order表的第一行记录", data)
	// 根据主键查询最后一条记录
	data, _ = dao.GetLastOrderInTable()
	fmt.Fprintln(w, "Order表的最后一条记录", data)
	// 使用查询用户的所有订单
	orders, _ := dao.GetUserOrders(534321467)
	fmt.Fprintln(w, "查询用户的订单, 条件 WHERE user_id = ?, 结果: ", orders)
	for _, order := range orders {
		fmt.Fprintln(w, "--", order.OrderNo)
	}

	userOrders, orderTotal, _ := logic.QueryUserOrdersInPage(534321467, 1, "2023-07-04", "")
	fmt.Fprintln(w, "分页查询用户的订单, 条件 WHERE user_id = ? AND created_at >= ? created, 返回结果: ", userOrders, "符合条件的总数: ", orderTotal)
	for _, order := range userOrders {
		fmt.Fprintln(w, "--", order.OrderNo, "创建时间", order.CreatedAt)
	}
	serialInfo, err := dao.GetOrderSerialInfo("20230704000000001366262146")
	fmt.Fprintln(w, "使用SELECT 指定查询返回的字段, 返回结果: ", serialInfo, "Error: ", err)
	serialInfoList, err := dao.GetOrdersSerial([]string{"20230704000000001366262146", "20230704000000001366372146"})
	fmt.Fprintln(w, "使用SELECT 指定查询返回的字段, 查询多条数据返回结果: ", serialInfoList, "Error: ", err)
	for _, serial := range serialInfoList {
		fmt.Fprintln(w, "--platform_order_no:", serial.OutOrderNo)
	}
	orderStat, _ := dao.GetOrderStat()
	//orderStat, _ := dao.GetOrderStatMap()
	fmt.Fprintln(w, "使用GROUP, 每日订单金额统计:", orderStat)
	for _, stat := range orderStat {
		fmt.Fprintln(w, "--Date:", stat.Date, "--Money:", stat.DayMoneyAmount)
	}

}

func SetOrderState(w http.ResponseWriter, r *http.Request) {
	err := dao.SetOrderStatePaySuccess("20230704000000001366262146")
	if err != nil {
		fmt.Fprintln(w, "更新失败!")
		return
	}
	fmt.Fprintln(w, "更新成功!")
	return
}

func MultipleTableQuery(w http.ResponseWriter, r *http.Request) {
	goodsList, err := dao.GetUserOrderGoods(534321467)
	if err != nil {
		fmt.Fprintln(w, "JOIN 查询", "Error: ", err)
		return
	}
	fmt.Fprintln(w, "JOIN 查询成功")
	for _, item := range goodsList {
		fmt.Fprintln(w, "--OrderNo:", item.OrderNo, "--GoodsName:", item.GoodsName, "---GoodsId:", item.GoodsId)
	}

	goodsMap, err := dao.GetUserOrderGoodsMap(534321467)
	fmt.Fprintln(w, "JOIN 查询", "Error: ", err)
	for key, item := range goodsMap {
		fmt.Fprintln(w, "--GoodsId:", key, "--GoodsName:", item)
	}
	return
}

func SetOrderPaySuccess(w http.ResponseWriter, r *http.Request) {
	err := logic.SetOrderSuccessAndCreateGoods(534321467, "20230704000000001366352146")
	if err != nil {
		fmt.Fprintln(w, "更新失败!")
		return
	}

	fmt.Fprintln(w, "更新成功!")
	return
}
