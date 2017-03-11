package models

import "github.com/revel/revel"

type Role struct {
	Role_id    int    `xorm:"pk"`
	Role_name  string `xorm:"varchar(11)"`
	Createdate int64  `xorm:"int"`
}

func (r *Role) Save(role Role) bool {
	has, err := DB_Write.Table("dl_sys_role").Insert(role)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (r *Role) Update(roleid int, role Role) bool {
	has, err := DB_Write.Table("dl_sys_role").Id(roleid).Update(role)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (r *Role) Getall() []Role {
	var role []Role
	err := DB_Write.Table("dl_sys_role").Find(&role)
	if err != nil {
		revel.WARN.Println("错误: %v", err)
		return nil
	}
	return role
}
