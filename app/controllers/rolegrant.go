package controllers

import "github.com/revel/revel"
import "drinking_life/app/models"
import "drinking_life/app/utils"
import "strconv"

type Rolegrant struct {
	*revel.Controller
}

func (r *Rolegrant) Save() revel.Result {
	model_grant := new(models.Rolegrant)
	role_id, _ := strconv.Atoi(r.Params.Get("roleid"))
	modelar_id, _ := strconv.Atoi(r.Params.Get("modularid"))
	var grant models.Rolegrant
	grant.Role_id = role_id
	grant.Modular_id = modelar_id
	grant.Createdate = utils.GetTimeNow()
	b := model_grant.Save(grant)
	if b == false {
		return r.RenderJson("角色授权失败")
	}
	return r.RenderJson("角色授权成功")
}
