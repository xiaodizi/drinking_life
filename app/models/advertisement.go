package models

import "github.com/revel/revel"
import "drinking_life/app/utils"

type Advertisement struct {
	Ad_id      int    `xorm:"pk"`
	Ad_name    string `xorm:"varchar(11)"`
	Photo      string `xorm:"varchar(255)"`
	Ad_words   string `xorm:"varchar(255)"`
	Is_display int    `xorm:"int(11)"`
	Sort_id    int    `xorm:"int(11)"`
	Class_id   int    `xorm:"int(11)"`
	Createdate int64  `xorm:"int"`
}

func (ad *Advertisement) Save(advertisement Advertisement) bool {
	has, err := DB_Write.Table("dl_ad").Insert(advertisement)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (ad *Advertisement) Delete(adid int) bool {
	advertisement := new(Advertisement)
	has, err := DB_Write.Table("dl_ad").Id(adid).Delete(advertisement)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (ad *Advertisement) Update(adid int, advertisement Advertisement) bool {
	has, err := DB_Write.Table("dl_ad").Id(adid).Update(advertisement)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (ad *Advertisement) Getall() []Advertisement {
	var advertisements []Advertisement
	err := DB_Write.Table("dl_ad").Find(&advertisements)
	if err != nil {
		revel.WARN.Println("错误: %v", err)
		return nil
	}
	return advertisements
}
func (ad *Advertisement) Sort(count int) []Advertisement {
	var advertisements []Advertisement
	err := DB_Write.Table("dl_ad").Limit(count).Asc("sort_id").Find(&advertisements)
	if err != nil {
		revel.WARN.Println("错误: %v", err)
		return nil
	}
	return advertisements
}
func (ad *Advertisement) IsDisplay(ad_id int) bool {
	has, err := DB_Write.Exec("update dl_ad set is_display = ? where ad_id = ?", utils.Is_display, ad_id)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
func (ad *Advertisement) IsnotDisplay(ad_id int) bool {
	has, err := DB_Write.Exec("update dl_ad set is_display = ? where ad_id = ?", utils.Is_not_display, ad_id)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
