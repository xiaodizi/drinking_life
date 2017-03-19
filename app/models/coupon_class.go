package models

import "github.com/revel/revel"

type CouponClass struct {
	Coupon_class_id   int    `xorm:"pk"`
	Coupon_class_name int    `xorm:"int(11)"`
	Product_id        string `xorm:"varchar(11)"`
}

func (cc *CouponClass) Save(couponClass CouponClass) bool {
	has, err := DB_Write.Table("dl_coupon_class").Insert(couponClass)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (cc *CouponClass) Delete(couponclassid int) bool {
	couponclass := new(CouponClass)
	has, err := DB_Write.Table("dl_coupon_class").Id(couponclassid).Delete(couponclass)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
func (cc *CouponClass) Update(couponclassid int, couponClass CouponClass) bool {
	has, err := DB_Write.Table("dl_coupon_class").Id(couponclassid).Update(couponClass)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
func (cc *CouponClass) Getall() []CouponClass {
	var couponClass []CouponClass
	err := DB_Write.Table("dl_coupon_class").Find(&couponClass)
	if err != nil {
		revel.WARN.Println("错误: %v", err)
		return nil
	}
	return couponClass
}
