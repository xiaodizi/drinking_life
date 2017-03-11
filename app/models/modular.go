package models

import "github.com/revel/revel"

type Modular struct {
	Modular_id   int    `xorm:"pk"`
	Modular_name string `xorm:"varchar(11)"`
	Modular_url  string `xorm:"varchar(256)"`
	Createdate   int64  `xorm:int`
}

func (m *Modular) Save(modular Modular) bool {
	has, err := DB_Write.Table("dl_sys_modular").Insert(modular)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (m *Modular) Update(modular Modular) bool {
	has, err := DB_Write.Table("dl_sys_modular").Exec("update set modular_name = ? ,modular_url = ? where modular_id = ? ", modular.Modular_name, modular.Modular_url, modular.Modular_url)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (m *Modular) Getall() []Modular {
	var modular []Modular
	err := DB_Read.Table("dl_sys_modular").Find(&modular)
	if err != nil {
		revel.WARN.Println("错误: %v", err)
		return nil
	}
	return modular
}
