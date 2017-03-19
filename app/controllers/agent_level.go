package controllers

import "github.com/revel/revel"
import "drinking_life/app/models"
import "strconv"

type Agentlevel struct {
	*revel.Controller
}

func (al *Agentlevel) Save() revel.Result {
	models_agent_level := new(models.Agentlevel)
	levelname := al.Params.Get("levelname")
	levelinfo := al.Params.Get("levelinfo")
	divided := al.Params.Get("divided")
	if levelname == "" {
		return al.RenderJson("代理商等级名称为空,添加失败")
	}
	if levelinfo == "" {
		return al.RenderJson("代理商等级描述为空,添加失败")
	}
	if divided == "" {
		return al.RenderJson("分成比例(百分比)为空,添加失败")
	}
	var agentlevel models.Agentlevel
	agentlevel.Level_name = levelname
	agentlevel.Level_info = levelinfo
	agentlevel.Divided = divided
	b := models_agent_level.Save(agentlevel)
	if b == false {
		return al.RenderJson("代理商等级添加失败")
	}
	return al.RenderJson("代理商等级添加成功")
}

func (al *Agentlevel) Delete() revel.Result {
	id, err := strconv.Atoi(al.Params.Get("id"))
	if err != nil {
		return al.RenderJson("代理商等级删除失败")
	}
	models_agent_level := new(models.Agentlevel)
	b := models_agent_level.Delete(id)
	if b == false {
		return al.RenderJson("代理商等级删除失败")
	}
	return al.RenderJson("代理商等级删除成功")
}
func (al *Agentlevel) Update() revel.Result {
	id, err := strconv.Atoi(al.Params.Get("id"))
	if err != nil {
		return al.RenderJson("代理商等级修改失败")
	}
	levelname := al.Params.Get("levelname")
	levelinfo := al.Params.Get("levelinfo")
	divided := al.Params.Get("divided")
	models_agent_level := new(models.Agentlevel)
	var agentlevel models.Agentlevel
	agentlevel.Level_name = levelname
	agentlevel.Level_info = levelinfo
	agentlevel.Divided = divided
	b := models_agent_level.Update(id, agentlevel)
	if b == false {
		return al.RenderJson("代理商等级修改失败")
	}
	return al.RenderJson("代理商等级修改成功")
}

func (al *Agentlevel) Getall() revel.Result {
	models_agent_level := new(models.Agentlevel)
	var agentlevels []models.Agentlevel
	agentlevels = models_agent_level.Getall()
	return al.RenderJson(agentlevels)
}
