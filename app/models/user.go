package models

import "github.com/revel/revel"
import "strconv"

type User struct {
	Userid      int64  `xorm:"pk"`
	Cellphone   string `xorm:"varchar(11)"`
	Wechat_t    string `xorm:"varchar(256)"`
	Wechat_name string `xorm:"varchar(11)"`
	Account_id  int64  `xorm:"int"`
	Balance     int64  `xorm:"int"`
	Createtime  int64  `xorm:"int"`
}

func (u *User) Getall() []User {
	var user []User
	err := DB_Read.Table("user").Find(&user)
	if err != nil {
		revel.WARN.Println("错误: %v", err)
	}
	return user
}

func (u *User) Getbyid(cellphone string, userid int) *User {
	user := new(User)
	where := "1=1"
	if userid > 0 {
		where = where + "and userid = " + strconv.Itoa(userid)
	}
	if cellphone != "" {
		where = where + " and cellphone = " + cellphone
	}

	has, err := DB_Read.Table("user").Where(where).Get(user)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
	}
	revel.WARN.Println("testing:" + where)
	return user
}

func (u *User) Update(userid int64, user *User) bool {
	has, err := DB_Write.Table("user").Id(userid).Update(user)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (u *User) Save(user User) bool {
	has, err := DB_Write.Table("user").Insert(user)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	return true
}

func (u *User) Delete(userid int64) bool {
	user := new(User)
	has, err := DB_Write.Table("user").Id(userid).Delete(user)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误:%v", err)
		return false
	}
	return true
}

func (u *User) Hascellphone(cellphone string) bool {
	user := new(User)
	counts, err := DB_Read.Table("user").Where("cellphone = ?", cellphone).Count(user)
	if err != nil {
		revel.WARN.Println("错误: %v", err)
		return false
	}
	if counts > 0 {
		return false
	}
	return true

}
