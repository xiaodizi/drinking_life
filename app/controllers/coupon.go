package controllers

import "github.com/revel/revel"
import "drinking_life/app/models"
import "drinking_life/app/utils"
import "strconv"

type Coupon struct {
	*revel.Controller
}

func (c *Coupon) Save() revel.Result {
	models_coupon := new(models.Coupon)
	couponclassid, _ := strconv.Atoi(c.Params.Get("couponclassid"))
	isuse, _ := strconv.Atoi(c.Params.Get("isuse"))
	amount, _ := strconv.Atoi(c.Params.Get("amount"))
	expiredate, _ := strconv.Atoi(c.Params.Get("expiredate"))
	useproductid := c.Params.Get("useproductid")
	if couponclassid == 0 {
		return c.RenderJson("优惠券分类为空,添加失败")
	}
	if amount == 0 {
		return c.RenderJson("金额为空,添加失败")
	}
	if expiredate == 0 {
		return c.RenderJson("到期日为空,添加失败")
	}
	if useproductid == "" {
		return c.RenderJson("可以使用的产品为空,添加失败")
	}

	var coupon models.Coupon
	coupon.Coupon_class_id = couponclassid
	coupon.Isuse = isuse
	coupon.Amount = amount
	coupon.Expiredate = expiredate
	coupon.Receivedate = utils.GetTimeNow()
	coupon.Use_productid = useproductid
	b := models_coupon.Save(coupon)
	if b == false {
		return c.RenderJson("优惠券添加失败")
	}
	return c.RenderJson("优惠券添加成功")
}

func (c *Coupon) Delete() revel.Result {
	couponid, err := strconv.Atoi(c.Params.Get("couponid"))
	if err != nil {
		return c.RenderJson("优惠券删除失败")
	}
	models_coupon := new(models.Coupon)
	b := models_coupon.Delete(couponid)
	if b == false {
		return c.RenderJson("优惠券删除失败")
	}
	return c.RenderJson("优惠券删除成功")
}

func (c *Coupon) Update() revel.Result {
	couponid, _ := strconv.Atoi(c.Params.Get("couponid"))
	couponclassid, _ := strconv.Atoi(c.Params.Get("couponclassid"))
	amount, _ := strconv.Atoi(c.Params.Get("amount"))
	expiredate, _ := strconv.Atoi(c.Params.Get("expiredate"))
	useproductid := c.Params.Get("useproductid")
	models_coupon := new(models.Coupon)
	var coupon models.Coupon
	coupon.Coupon_class_id = couponclassid
	coupon.Amount = amount
	coupon.Expiredate = expiredate
	coupon.Use_productid = useproductid
	b := models_coupon.Update(couponid, coupon)
	if b == false {
		return c.RenderJson("优惠券修改失败")
	}
	return c.RenderJson("优惠券修改成功")
}

func (c *Coupon) Getall() revel.Result {
	models_coupon := new(models.Coupon)
	var coupons []models.Coupon
	coupons = models_coupon.Getall()
	return c.RenderJson(coupons)
}

func (c *Coupon) Isuse() revel.Result {
	couponid, _ := strconv.Atoi(c.Params.Get("couponid"))
	model_coupon := new(models.Coupon)
	b := model_coupon.IsUse(couponid)
	if b == false {
		return c.RenderJson("优惠券使用失败")
	}
	return c.RenderJson("优惠券使用成功")
}
func (c *Coupon) Isnotuse() revel.Result {
	couponid, _ := strconv.Atoi(c.Params.Get("couponid"))
	model_coupon := new(models.Coupon)
	b := model_coupon.IsnotUse(couponid)
	if b == false {
		return c.RenderJson("优惠券使用失败")
	}
	return c.RenderJson("优惠券使用成功")
}
