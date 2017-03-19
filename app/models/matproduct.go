package main

import "github.com/revel/revel"

type Matproduct struct {
	Id             int    `xorm:"pk"`
	Name           string `xorm:"VARCHAR(20)"`
	Price          int64  `xorm:"int"`
	Prod_date      int64  `xorm:"int"`
	Level          int    `xorm:"int"`
	Prod_company   string `xorm:"VARCHAR(30)"`
	Prod_contact   string `xorm:"VARCHAR(5)"`
	Prod_cellphone int    `xorm:"int"`
	Prod_phone     string `xorm:"VARCHAR(20)"`
	Pubdate        int64  `xorm:"int"`
}

func (m *Matproduct) Save(product Matproduct) bool {
	has, err := DB_Write.Table("dl_material_product").Insert(product)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
func (m *Matproduct) Getall() []Mapproduct {
	var product []Matproduct
	err := DB_Wirite.Table("dl_material_product").Find(&product)
	if err != nil {
		revel.WARN.Println("错误: %v", err)
		return nil
	}
	return product
}

func (m *Matproduct) Getbyid(id int) Matproduct {
	product := new(Matproduct)
	has, err := DB_Write.Table("dl_material_product").Where("id = ?", id).Get(product)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误： %v", err)
		return nil
	}
	return product
}

func (m *Matproduct) Update(id int, product Matproduct) bool {
	has, err := DB_Wirite.Table("dl_material_product").Id(id).Update(product)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
