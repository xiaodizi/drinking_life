package controllers

import "github.com/revel/revel"
import "drinking_life/app/models"
import "strconv"

type CouponClass struct {
	*revel.Controller
}

func (cc *CouponClass) Save() revel.Result {
	models_coupon_class := new(models.CouponClass)
	couponclassname, _ := strconv.Atoi(cc.Params.Get("couponclassname"))
	productid := cc.Params.Get("productid")
	if couponclassname == 0 {
		return cc.RenderJson("优惠券类型名称为空,添加失败")
	}
	if productid == "" {
		return cc.RenderJson("限制购买产品为空,添加失败")
	}
	var couponclass models.CouponClass
	couponclass.Coupon_class_name = couponclassname
	couponclass.Product_id = productid
	b := models_coupon_class.Save(couponclass)
	if b == false {
		return cc.RenderJson("优惠券类型添加失败")
	}
	return cc.RenderJson("优惠券类型添加成功")
}

func (cc *CouponClass) Delete() revel.Result {
	couponclassid, err := strconv.Atoi(cc.Params.Get("couponclassid"))
	if err != nil {
		return cc.RenderJson("优惠券类型删除失败")
	}
	models_coupon_class := new(models.CouponClass)
	b := models_coupon_class.Delete(couponclassid)
	if b == false {
		return cc.RenderJson("优惠券类型删除失败")
	}
	return cc.RenderJson("优惠券类型删除成功")
}
func (cc *CouponClass) Update() revel.Result {
	couponclassid, err := strconv.Atoi(cc.Params.Get("couponclassid"))
	if err != nil {
		return cc.RenderJson("优惠券类型修改失败")
	}
	couponclassname, _ := strconv.Atoi(cc.Params.Get("couponclassname"))
	productid := cc.Params.Get("productid")
	models_coupon_class := new(models.CouponClass)
	var couponclasses models.CouponClass
	couponclasses.Coupon_class_name = couponclassname
	couponclasses.Product_id = productid
	b := models_coupon_class.Update(couponclassid, couponclasses)
	if b == false {
		return cc.RenderJson("类型修改失败")
	}
	return cc.RenderJson("类型修改成功")
}

func (cc *CouponClass) Getall() revel.Result {
	model_coupon_class := new(models.CouponClass)
	var couponclasses []models.CouponClass
	couponclasses = model_coupon_class.Getall()
	return cc.RenderJson(couponclasses)
}
