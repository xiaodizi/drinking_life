package controllers

import "drinking_life/app/models"
import "github.com/revel/revel"
import "drinking_life/app/utils"

type User struct {
	*revel.Controller
}

func (u *User) Userlist() revel.Result {
	model_user := new(models.User)
	userlist := model_user.Getall()
	return u.RenderJson(userlist)
}

func (u *User) Save() revel.Result {
	var user models.User
	user.Balance = 0
	user.Wechat_t = u.Params.Get("webchat_t")
	user.Cellphone = u.Params.Get("cellphone")
	user.Wechat_name = u.Params.Get("wechat_name")
	user.Createtime = utils.GetTimeNow()
	models_user := new(models.User)
	b := models_user.Save(user)
	if b == true {
		u := models_user.Getbyid(user.Cellphone, 0)
		account := new(models.Account)
		var acc models.Account
		acc.User_id = u.Userid
		bo := account.Save(acc)
		if bo == false {
			revel.WARN.Println("错误: account数据添加失败")
			o := models_user.Delete(u.Userid)
			if o == true {
				revel.WARN.Println("数据回滚成功,已删除")
			}
		}
	} else {
		revel.WARN.Println("错误:会员注册失败")
	}
	return u.RenderJson(b)
}

func (u *User) Hascellphone() revel.Result {
	cellphone := u.Params.Get("cellphone")
	model_user := new(models.User)
	b := model_user.Hascellphone(cellphone)
	if b == true {
		return u.RenderJson("未注册")
	} else {
		return u.RenderJson("已注册")
	}
}
