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
		return c.RenderJson("price数据格式错误: " + err3.Error())
	}
	userId, err4 := strconv.Atoi(c.Params.Get("user_id"))
	if err4 != nil {
		revel.WARN.Println(err4)
		return c.RenderJson("user_id数据格式错误:" + err4.Error())
	}
	cart.User_id = userId
	price := int(p * 100)
	cart.Cart_price = price
	num, err2 := strconv.Atoi(c.Params.Get("num"))
	if err2 != nil {
		revel.WARN.Println(err2)
		return c.RenderJson("num数据格式错误:" + err2.Error())
	}
	cart.Cart_pro_num = num
	cart.Cart_pro_name = c.Params.Get("pro_name")
	pro_id, err1 := strconv.Atoi(c.Params.Get("pro_id"))
	if err1 != nil {
		revel.WARN.Println(err3)
		return c.RenderJson("pro_id数据格式错误:" + err1.Error())
	}
	cart.Cart_pro_id = pro_id
	ct := new(models.Cart)
	b := ct.Save(cart)
	if b == false {
		return c.RenderJson("购物车添加失败!")
	}
	return c.RenderJson("购物车添加成功")
}

func (c *Cart) Addnum() revel.Result {
	ct := new(models.Cart)
	cartNum, _ := strconv.Atoi(c.Params.Get("num"))
	cartId, _ := strconv.Atoi(c.Params.Get("cartid"))
	carts := ct.GetByCartId(cartId)
	if cartNum == 0 {
		return c.RenderJson("数量不能为空")
	}
	var cartes models.Cart
	cartes.Cart_pro_name = carts.Cart_pro_name
	cartes.Cart_pro_num = cartNum
	cartes.Cart_price = carts.Cart_price * cartNum
	cartes.Cart_pro_id = carts.Cart_pro_id
	cartes.Creattime = utils.GetTimeNow()
	cartes.User_id = carts.User_id
	b := ct.Upadte(cartId, cartes)
	if b == false {
		return c.RenderJson("数量修改失败")
	}
	return c.RenderJson("数量修改成功")
}

func (c *Cart) Delete() revel.Result {
	ct := new(models.Cart)
	cartId, _ := strconv.Atoi(c.Params.Get("cartId"))
	b := ct.Delete(cartId)
	if b == false {
		return c.RenderJson("购物车删除失败")
	}
	return c.RenderJson("购物车删除成功")
}
