package controllers

import (
	"regexp"

	"github.com/revel/revel"
)
import "drinking_life/app/models"
import "strconv"
import "drinking_life/app/utils"

type Agent struct {
	*revel.Controller
}

func validate(mobileNum string) bool {
	reg := regexp.MustCompile(utils.Regular)
	return reg.MatchString(mobileNum)
}
func (a *Agent) Save() revel.Result {
	models_agent := new(models.Agent)
	name := a.Params.Get("name")
	cellphone, _ := strconv.Atoi(a.Params.Get("cellphone"))
	city, _ := strconv.Atoi(a.Params.Get("city"))
	address := a.Params.Get("address")
	bankno, _ := strconv.Atoi(a.Params.Get("bankno"))
	bank := a.Params.Get("bank")
	signing_date, _ := strconv.Atoi(a.Params.Get("signing_date"))
	expire_date, _ := strconv.Atoi(a.Params.Get("expire_date"))
	level, _ := strconv.Atoi(a.Params.Get("level"))
	if name == "" {
		return a.RenderJson("代理商名称为空,添加失败")
	}
	if !validate(a.Params.Get("cellphone")) {
		return a.RenderJson("代理商手机号格式不正确,添加失败")
	}
	if city == 0 {
		return a.RenderJson("代理商所在城市为空,添加失败")
	}
	if address == "" {
		return a.RenderJson("代理商地址为空,添加失败")
	}
	if bankno == 0 {
		return a.RenderJson("银行账号为空,添加失败")
	}
	if bank == "" {
		return a.RenderJson("所在银行为空,添加失败")
	}
	if signing_date == 0 {
		return a.RenderJson("签约日期为空,添加失败")
	}
	if expire_date == 0 {
		return a.RenderJson("合同到期日为空,添加失败")
	}
	if level == 0 {
		return a.RenderJson("代理商等级为空,添加失败")
	}

	var agent models.Agent
	agent.Name = name
	agent.Cellphone = cellphone
	agent.City = city
	agent.Address = address
	agent.Bankno = bankno
	agent.Bank = bank
	agent.Signing_date = signing_date
	agent.Expire_date = expire_date
	agent.Level = level
	b := models_agent.Save(agent)
	if b == false {
		return a.RenderJson("代理商添加失败")
	}
	return a.RenderJson("代理商添加成功")
}

func (a *Agent) Delete() revel.Result {
	id, err := strconv.Atoi(a.Params.Get("id"))
	if err != nil {
		return a.RenderJson("代理商删除失败")
	}
	models_agent := new(models.Agent)
	b := models_agent.Delete(id)
	if b == false {
		return a.RenderJson("代理商删除失败")
	}
	return a.RenderJson("代理商删除成功")
}
func (a *Agent) Update() revel.Result {
	id, err := strconv.Atoi(a.Params.Get("id"))
	if err != nil {
		return a.RenderJson("代理商修改失败")
	}
	name := a.Params.Get("name")
	cellphone, _ := strconv.Atoi(a.Params.Get("cellphone"))
	if !validate(a.Params.Get("cellphone")) {
		return a.RenderJson("代理商手机号格式不正确,添加失败")
	}
	city, _ := strconv.Atoi(a.Params.Get("city"))
	address := a.Params.Get("address")
	bankno, _ := strconv.Atoi(a.Params.Get("bankno"))
	bank := a.Params.Get("bank")
	signing_date, _ := strconv.Atoi(a.Params.Get("signing_date"))
	expire_date, _ := strconv.Atoi(a.Params.Get("expire_date"))
	level, _ := strconv.Atoi(a.Params.Get("level"))
	models_agent := new(models.Agent)
	var agent models.Agent
	agent.Name = name
	agent.Cellphone = cellphone
	agent.City = city
	agent.Address = address
	agent.Bankno = bankno
	agent.Bank = bank
	agent.Signing_date = signing_date
	agent.Expire_date = expire_date
	agent.Level = level
	b := models_agent.Update(id, agent)
	if b == false {
		return a.RenderJson("代理商修改失败")
	}
	return a.RenderJson("代理商修改成功")
}

func (a *Agent) Getall() revel.Result {
	models_agent := new(models.Agent)
	var agents []models.Agent
	agents = models_agent.Getall()
	return a.RenderJson(agents)
}
