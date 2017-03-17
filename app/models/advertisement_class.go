package models

import "github.com/revel/revel"

type AdClass struct {
	Class_id   int    `xorm:"pk"`
	Class_name string `xorm:"varchar(11)"`
	Createdate int64  `xorm:"int"`
}

func (ac *AdClass) Save(adclass AdClass) bool {
	has, err := DB_Write.Table("dl_ad_class").Insert(adclass)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (ac *AdClass) Delete(classid int) bool {
	adclass := new(AdClass)
	has, err := DB_Write.Table("dl_ad_class").Id(classid).Delete(adclass)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
func (ac *AdClass) Update(classid int, adclass AdClass) bool {
	has, err := DB_Write.Table("dl_ad_class").Id(classid).Update(adclass)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
func (ac *AdClass) Getall() []AdClass {
	var adclass []AdClass
	err := DB_Write.Table("dl_ad_class").Find(&adclass)
	if err != nil {
		revel.WARN.Println("错误: %v", err)
		return nil
	}
	return adclass
}
