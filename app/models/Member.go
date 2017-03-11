package models

import "github.com/revel/revel"

type Member struct {
	Member_id         int    `xorm:"pk"`
	Member_name       string `xorm:"varchar(11)"`
	Member_password   string `xorm:"varchar(256)"`
	Member_role_id    int    `xorm:"int"`
	Member_cellphone  int    `xorm:"varchar(20)"`
	Member_createdate int64  `xorm:"int"`
	Is_use            int    `xorm:"int"`
}

func (m *Member) Save(member Member) bool {
	has, err := DB_Write.Table("dl_sys_member").Insert(member)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	if has == 0 {
		return false
	}
	return true
}

func (m *Member) Updatepassword(newPassword string, memberid int) bool {
	has, err := DB_Write.Table("dl_sys_member").Exec("update dl_sys_member set member_password = ? where member_id = ?", newPassword, memberid)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误:%v", err)
		return false
	}
	return true
}

func (m *Member) Isuse(isUse int, memberid int) bool {
	has, err := DB_Write.Table("dl_sys_member").Exec("update dl_sys_member set is_use = ? where member_id = ?", isUse, memberid)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v ", err)
		return false
	}
	return true
}

func (m *Member) Getoldpass(memberid int) string {
	member := new(Member)
	has, err := DB_Read.Table("dl_sys_member").Select("member_id,member_password").Where("member_id = ?", memberid).Get(member)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误:%v", err)
	}
	return member.Member_password
}

func (m *Member) Hasname(name string, cellphone string) bool {
	member := new(Member)
	where := "1=1"
	if name != "" {
		where += " and member_name = '" + name + "'"
	}
	if cellphone != "" {
		where += " and member_cellphone = '" + cellphone + "'"
	}
	has, err := DB_Read.Table("dl_sys_member").Where(where).Count(member)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}
	if has > 0 {
		return false
	}
	return true
}

func (m *Member) Vaildname(username string, password string) bool {
	member := new(Member)
	has, err := DB_Read.Table("dl_sys_member").Where("member_name ＝'" + username + "' and member_password = '" + password + "'").Count(member)
	if err != nil {
		revel.WARN.Println("错误: %v", err)
		revel.WARN.Println(has)
		return false
	}
	if has < 0 {
		return false
	}
	return true
}
