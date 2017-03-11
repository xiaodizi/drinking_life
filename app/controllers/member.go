package controllers

import "github.com/revel/revel"
import "drinking_life/app/utils"
import "drinking_life/app/models"
import "strconv"

type Member struct {
	*revel.Controller
}

func (m *Member) Save() revel.Result {
	var member models.Member
	member.Member_name = m.Params.Get("name")
	member.Member_role_id, _ = strconv.Atoi(m.Params.Get("roleid"))
	member.Member_cellphone, _ = strconv.Atoi(m.Params.Get("cellphone"))
	member.Member_password = utils.HashMd5("123456")
	member.Member_createdate = utils.GetTimeNow()
	member.Is_use = utils.Is_not_use

	model_member := new(models.Member)
	b := model_member.Save(member)
	if b == false {
		return m.RenderJson("添加失败")
	}
	return m.RenderJson("添加成功")
}

func (u *Member) Update() revel.Result {
	member_id, _ := strconv.Atoi(u.Params.Get("member_id"))
	oldPassword := u.Params.Get("oldpass")
	newPassword := utils.HashMd5(u.Params.Get("newpass"))
	model_member := new(models.Member)
	oldpass := model_member.Getoldpass(member_id)
	revel.WARN.Println("testing:" + oldpass)
	if oldpass != utils.HashMd5(oldPassword) {
		return u.RenderJson("旧密码输入错误")
	} else {
		b := model_member.Updatepassword(newPassword, member_id)
		if b == false {
			return u.RenderJson("密码修改失败")
		}
	}
	return u.RenderJson("密码修改成功")
}

func (u *Member) Hasbyname() revel.Result {
	cellphone := u.Params.Get("cellphone")
	name := u.Params.Get("name")
	model_member := new(models.Member)
	b := model_member.Hasname(name, cellphone)
	if b == false {
		return u.RenderJson("此手机号或者姓名已添加")
	}
	return u.RenderJson("未添加手机号或者姓名")
}

func (u *Member) Hascode() revel.Result {
	hasCode := utils.Krand(6, utils.KC_RAND_KIND_ALL)
	u.Session["hascode"] = hasCode
	return u.RenderJson("验证码:" + hasCode)
}

func (u *Member) Login() revel.Result {
	if u.Request.Method == "GET" {

	} else {
		userName := u.Params.Get("username")
		password := u.Params.Get("password")
		hascode := u.Params.Get("hascode")
		if hascode != utils.GetSession("hascode", u.Session) {
			revel.WARN.Println("testing:" + hascode)
			return u.RenderJson("验证码输入错误")
		}
		models_member := new(models.Member)
		b := models_member.Vaildname(userName, password)
		revel.WARN.Println("string:" + userName)
		if b == false {
			return u.RenderJson("登陆失败,用户名或者密码错误")
		}
	}
	return u.RenderJson("登陆成功")
}

func (u *Member) Isuse() revel.Result {
	memberid, _ := strconv.Atoi(u.Params.Get("member"))
	isuse, _ := strconv.Atoi(u.Params.Get("isuse"))
	model_member := new(models.Member)
	b := model_member.Isuse(isuse, memberid)
	if b == false {
		return u.RenderJson("修改状态失败")
	}
	return u.RenderJson("修改状态成功")
}
