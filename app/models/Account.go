package models

import "github.com/revel/revel"

type Account struct {
	Account_id int64 `xorm:"pk"`
	User_id    int64 `xorm:"int"`
}

func (a *Account) Save(account Account) bool {
	has, err := DB_Write.Table("dl_sys_account").Insert(account)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
