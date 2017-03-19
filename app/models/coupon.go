package models

import "github.com/revel/revel"
import "drinking_life/app/utils"

type Coupon struct {
	Coupon_id       int    `xorm:"pk"`
	Coupon_class_id int    `xorm:"int(11)"`
	Isuse           int    `xorm:"int(255)"`
	Amount          int    `xorm:"int(255)"`
	Expiredate      int    `xorm:"int(11)"`
	Receivedate     int64  `xorm:"int(11)"`
	Use_productid   string `xorm:"varchar(11)"`
}

func (c *Coupon) Save(coupon Coupon) bool {
	has, err := DB_Write.Table("dl_coupon").Insert(coupon)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (c *Coupon) Delete(couponid int) bool {
	coupon := new(Coupon)
	has, err := DB_Write.Table("dl_coupon").Id(couponid).Delete(coupon)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (c *Coupon) Update(couponid int, coupon Coupon) bool {
	has, err := DB_Write.Table("dl_coupon").Id(couponid).Update(coupon)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (c *Coupon) Getall() []Coupon {
	var coupons []Coupon
	err := DB_Write.Table("dl_coupon").Find(&coupons)
	if err != nil {
		revel.WARN.Println("错误: %v", err)
		return nil
	}
	return coupons
}
func (c *Coupon) IsUse(coupon_id int) bool {
	has, err := DB_Write.Exec("update dl_coupon set isUse = ? where coupon_id = ?", utils.Is_use, coupon_id)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
func (c *Coupon) IsnotUse(coupon_id int) bool {
	has, err := DB_Write.Exec("update dl_coupon set isUse = ? where coupon_id = ?", utils.Is_not_use, coupon_id)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
