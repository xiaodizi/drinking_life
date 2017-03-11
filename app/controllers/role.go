package controllers

import "github.com/revel/revel"
import "drinking_life/app/models"
import "drinking_life/app/utils"
import "strconv"

type Role struct {
	*revel.Controller
}

func (r *Role) Save() revel.Result {
	models_role := new(models.Role)
	rolename := r.Params.Get("rolename")
	if rolename == "" {
		return r.RenderJson("角色名称为空,添加失败")
	}
	var role models.Role
	role.Role_name = rolename
	role.Createdate = utils.GetTimeNow()
	b := models_role.Save(role)
	if b == false {
		return r.RenderJson("角色添加失败")
	}
	return r.RenderJson("角色添加成功")
}

func (r *Role) Update() revel.Result {
	roleid, _ := strconv.Atoi(r.Params.Get("roleid"))
	rolename := r.Params.Get("rolename")
	model_role := new(models.Role)
	var roles models.Role
	roles.Role_name = rolename
	b := model_role.Update(roleid, roles)
	if b == false {
		return r.RenderJson("角色修改失败")
	}
	return r.RenderJson("角色修改成功")
}

func (r *Role) Getall() revel.Result {
	model_role := new(models.Role)
	var roles []models.Role
	roles = model_role.Getall()
	return r.RenderJson(roles)
}
