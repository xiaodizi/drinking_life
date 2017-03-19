package controllers

import "github.com/revel/revel"
import "drinking_life/app/models"
import "strconv"

type Equipment struct {
	*revel.Controller
}

func (e *Equipment) Save() revel.Result {
	models_equipment := new(models.Equipment)
	eq_no := e.Params.Get("eqno")
	eq_province, _ := strconv.Atoi(e.Params.Get("eqprovince"))
	eq_city, _ := strconv.Atoi(e.Params.Get("eqcity"))
	eq_county, _ := strconv.Atoi(e.Params.Get("eqcounty"))
	agent_id, _ := strconv.Atoi(e.Params.Get("agentid"))
	if eq_no == "" {
		return e.RenderJson("设备序列号为空,添加失败")
	}
	if eq_province == 0 {
		return e.RenderJson("设备所在省为空,添加失败")
	}
	if eq_city == 0 {
		return e.RenderJson("设备所在市为空,添加失败")
	}
	if eq_county == 0 {
		return e.RenderJson("设备所在县为空,添加失败")
	}
	if agent_id == 0 {
		return e.RenderJson("直营代理商为空,添加失败")
	}
	var equipment models.Equipment
	equipment.Eq_no = eq_no
	equipment.Eq_province = eq_province
	equipment.Eq_city = eq_city
	equipment.Eq_county = eq_county
	equipment.Agent_id = agent_id
	b := models_equipment.Save(equipment)
	if b == false {
		return e.RenderJson("设备添加失败")
	}
	return e.RenderJson("设备添加成功")
}

func (e *Equipment) Delete() revel.Result {
	id, err := strconv.Atoi(e.Params.Get("id"))
	if err != nil {
		return e.RenderJson("设备删除失败")
	}
	models_equipment := new(models.Equipment)
	b := models_equipment.Delete(id)
	if b == false {
		return e.RenderJson("设备删除失败")
	}
	return e.RenderJson("设备删除成功")
}

func (e *Equipment) Update() revel.Result {
	id, _ := strconv.Atoi(e.Params.Get("id"))
	eq_no := e.Params.Get("eqno")
	eq_province, _ := strconv.Atoi(e.Params.Get("eqprovince"))
	eq_city, _ := strconv.Atoi(e.Params.Get("eqcity"))
	eq_county, _ := strconv.Atoi(e.Params.Get("eqcounty"))
	agent_id, _ := strconv.Atoi(e.Params.Get("agentid"))
	models_equipment := new(models.Equipment)
	var equipment models.Equipment
	equipment.Eq_no = eq_no
	equipment.Eq_province = eq_province
	equipment.Eq_city = eq_city
	equipment.Eq_county = eq_county
	equipment.Agent_id = agent_id
	b := models_equipment.Update(id, equipment)
	if b == false {
		return e.RenderJson("设备修改失败")
	}
	return e.RenderJson("设备修改成功")
}

func (e *Equipment) Getall() revel.Result {
	models_equipment := new(models.Equipment)
	var equipments []models.Equipment
	equipments = models_equipment.Getall()
	return e.RenderJson(equipments)
}
