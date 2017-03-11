package models

import "github.com/revel/revel"
import "drinking_life/app/utils"

type Product struct {
	Productid   int    `xorm:"pk"`
	Productname string `xorm:"varchar(11)"`
	Price       int    `xorm:"int"`
	Isframe     int    `xorm:"int"`
	Sort_id     int    `xorm:"int"`
	Createtime  int64  `xorm:"int"`
}

func (p *Product) Save(product Product) bool {
	has, err := DB_Write.Table("dl_product").Insert(product)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (p *Product) Update(product Product) bool {
	has, err := DB_Write.Table("dl_product").Exec("update set productname = ? , price = ?, sort_id = ?  where productid = ? ", product.Productname, product.Price, product.Isframe, product.Isframe, product.Sort_id, product.Productid)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (p *Product) Upframe(productid int) bool {
	has, err := DB_Write.Table("dl_product").Exec("update set isframe = ? where productid = ?", utils.Is_frame_up, productid)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (p *Product) Downframe(productid int) bool {
	has, err := DB_Write.Table("dl_product").Exec("update set isframe = ? where productid = ?", utils.Is_frame_down, productid)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v ", err)
		return false
	}
	return true
}

func (p *Product) Getall() []Product {
	var product []Product
	err := DB_Read.Table("dl_product").Find(&product)
	if err != nil {
		revel.WARN.Println("错误: %v", err)
	}
	return product
}
