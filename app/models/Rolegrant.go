package models

import "github.com/revel/revel"
import "strconv"

type Rolegrant struct {
	Grant_id   int   `xorm:"pk"`
	Role_id    int   `xorm:"int"`
	Modular_id int   `xorm:"int"`
	Createdate int64 `xorm:"int"`
}

func (r *Rolegrant) Save(grant Rolegrant) bool {
	has, err := DB_Write.Table("dl_sys_rolegrant").Insert(grant)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (r *Rolegrant) Delete(grand Rolegrant) bool {
	where := "1=1 "
	if grand.Grant_id != 0 {
		where += " and grant_id = " + strconv.Itoa(grand.Grant_id)
	}
	if grand.Role_id != 0 {
		where += " and role_id = " + strconv.Itoa(grand.Role_id)
	}
	if grand.Modular_id != 0 {
		where += " and modular_id = " + strconv.Itoa(grand.Modular_id)
	}
	rolegrant := new(Rolegrant)
	has, err := DB_Write.Table("dl_sys_rolegrant").Where(where).Delete(rolegrant)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
