package models

import "github.com/revel/revel"

type Cart struct {
	Cart_id       int    `xorm:"pk"`
	Cart_pro_name string `xorm:"varchar(11)"`
	Cart_pro_num  int    `xorm:"int"`
	Cart_price    int    `xorm:"int"`
	Creattime     int64  `xorm:"int"`
}
