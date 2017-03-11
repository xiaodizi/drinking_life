package controllers

import "github.com/revel/revel"
import "drinking_life/app/models"
import "drinking_life/app/utils"

type Modular struct {
	*revel.Controller
}

func (m *Modular) Save() revel.Result {
	modular_name := m.Params.Get("modular_name")
	modular_url := m.Params.Get("modular_url")
	var modular models.Modular
	modular.Modular_name = modular_name
	modular.Modular_url = modular_url
	modular.Createdate = utils.GetTimeNow()
	modulars := new(models.Modular)
	b := modulars.Save(modular)
	if b == false {
		return m.RenderJson("添加模块失败")
	}
	return m.RenderJson("添加模块成功")
}

func (m *Modular) Update() revel.Result {
	modular_name := m.Params.Get("modualr_name")
	modular_url := m.Params.Get("modular_url")
	var modular models.Modular
	modular.Modular_name = modular_name
	modular.Modular_url = modular_url
	md := new(models.Modular)
	b := md.Save(modular)
	if b == false {
		return m.RenderJson("模块修改失败")
	}
	return m.RenderJson("模块修改成功")
}

func (m *Modular) Getall() revel.Result {
	model_modular := new(models.Modular)
	modular := model_modular.Getall()
	return m.RenderJson(modular)
}
