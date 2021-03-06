package models

import "github.com/revel/revel"
import "strconv"

type Cart struct {
	Cart_id       int    `xorm:"pk"`
	Cart_pro_name string `xorm:"varchar(11)"`
	Cart_pro_num  int    `xorm:"int"`
	Cart_price    int    `xorm:"int"`
	Cart_pro_id   int    `xorm:"int"`
	Creattime     int64  `xorm:"int"`
	User_id       int    `xorm:"int"`
}

func (c *Cart) Save(cart Cart) bool {
	has, err := DB_Write.Table("dl_cart").Insert(cart)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (c *Cart) Delete(cartId int) bool {
	cart := new(Cart)
	has, err := DB_Write.Table("dl_cart").Where("cart_id").Delete(cart)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (c *Cart) Upadte(cartId int, cart Cart) bool {
	has, err := DB_Write.Table("dl_cart").Id(cartId).Update(cart)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (c *Cart) GetByCartId(cartId int) *Cart {
	cart := new(Cart)
	has, err := DB_Read.Table("dl_cart").Where("cart_id = '" + strconv.Itoa(cartId) + "'").Get(cart)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println(err.Error())
	}
	return cart
}
