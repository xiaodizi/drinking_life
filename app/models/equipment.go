package models

import "github.com/revel/revel"

type Equipment struct {
	Id          int    `xorm:"pk"`
	Eq_no       string `xorm:"varchar(11)"`
	Eq_province int    `xorm:"int(11)"`
	Eq_city     int    `xorm:"int(11)"`
	Eq_county   int    `xorm:"int(11)"`
	Agent_id    int    `xorm:"int(11)"`
}

func (e *Equipment) Save(equipment Equipment) bool {
	has, err := DB_Write.Table("dl_equipment").Insert(equipment)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (e *Equipment) Delete(id int) bool {
	equipment := new(Equipment)
	has, err := DB_Write.Table("dl_equipment").Id(id).Delete(equipment)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (e *Equipment) Update(id int, equipment Equipment) bool {
	has, err := DB_Write.Table("dl_equipment").Id(id).Update(equipment)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (e *Equipment) Getall() []Equipment {
	var equipments []Equipment
	err := DB_Write.Table("dl_equipment").Find(&equipments)
	if err != nil {
		revel.WARN.Println("错误: %v", err)
		return nil
	}
	return equipments
}
