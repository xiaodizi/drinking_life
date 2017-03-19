package models

import "github.com/revel/revel"

type AccountRecord struct {
	Record_id   int    `xorm:"pk"`
	Reason      int    `xorm:"int(11)"`
	Amount      int    `xorm:"int(11)"`
	Pubdate     int64  `xorm:"int(11)"`
	Type        int    `xorm:"int(11)"`
	Transfer_no string `xorm:"varchar(11)"`
	Account_id  int    `xorm:"int(11)"`
}

func (ar *AccountRecord) Save(accountrecord AccountRecord) bool {
	has, err := DB_Write.Table("account_record").Insert(accountrecord)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (ar *AccountRecord) Delete(recordid int) bool {
	accountrecord := new(AccountRecord)
	has, err := DB_Write.Table("account_record").Id(recordid).Delete(accountrecord)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
func (ar *AccountRecord) Update(recordid int, accountrecord AccountRecord) bool {
	has, err := DB_Write.Table("account_record").Id(recordid).Update(accountrecord)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}
func (ar *AccountRecord) Getall() []AccountRecord {
	var accountrecords []AccountRecord
	err := DB_Write.Table("account_record").Find(&accountrecords)
	if err != nil {
		revel.WARN.Println("错误: %v", err)
		return nil
	}
	return accountrecords
}
