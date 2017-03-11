package models

import "github.com/revel/revel"

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
