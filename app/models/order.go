package models

import "github.com/revel/revel"
import "drinking_life/app/utils"

type Order struct {
	Orderid     int64
	Ispayment   int
	Isreceive   int
	City        int
	Productid   int
	Productname string
	Price       int
	Userid      int64
	Eq_id       int
	Agent_id    int64
	Pay_time    int64
	Finish_time int64
	Create_time int64
}

func (o *Order) Save(order Order) bool {
	has, err := DB_Write.Table("dl_order").Insert(order)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (o *Order) Payment(orderId int64) bool {
	has, err := DB_Write.Table("dl_order").Exec("update dl_order set ispayment = ?,pay_time = ? where orderid = ?", utils.ALERD_PAYID, utils.GetTimeNow(), orderId)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (o *Order) Alreadreceive(orderId int64) bool {
	has, err := DB_Write.Table("dl_order").Exec("update dl_order set isreceive = ?,finish_time = ? where orderid = ? ", utils.ALERD_RECEIVE, utils.GetTimeNow, orderId)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
