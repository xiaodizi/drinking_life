package models

import "github.com/revel/revel"

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
