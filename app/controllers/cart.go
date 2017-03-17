package controllers

import "github.com/revel/revel"
import "drinking_life/app/models"
import "drinking_life/app/utils"
import "strconv"

type Cart struct {
	*revel.Controller
}

func (c *Cart) Save() revel.Result {
	var cart models.Cart
	cart.Creattime = utils.GetTimeNow()
	p, err3 := strconv.ParseFloat(c.Params.Get("price"), 32)
	if err3 != nil {
		revel.WARN.Println(err3)
		return c.RenderJson(err3)
	}
	price := int(p * 100)
	cart.Cart_price = price
	num, err2 := strconv.Atoi(c.Params.Get("num"))
	if err2 != nil {
		revel.WARN.Println(err2)
		return c.RenderJson(err2)
	}
	cart.Cart_pro_num = num
	cart.Cart_pro_name = c.Params.Get("pro_name")
	pro_id, err3 := strconv.Atoi(c.Params.Get("pro_id"))
	if err3 != nil {
		revel.WARN.Println(err3)
		return c.RenderJson(err3)
	}
	cart.Cart_pro_id = pro_id
	ct := new(models.Cart)
	b := ct.Save(cart)
	if b == false {
		return c.RenderJson("购物车添加失败!")
	}
	return c.RenderJson("购物车添加成功")
}
