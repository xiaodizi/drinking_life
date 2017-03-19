package controllers

import "github.com/revel/revel"
import "drinking_life/app/models"
import "drinking_life/app/utils"
import "strconv"

type AdClass struct {
	*revel.Controller
}

func (ac *AdClass) Save() revel.Result {
	models_ad_class := new(models.AdClass)
	classname := ac.Params.Get("classname")
	if classname == "" {
		return ac.RenderJson("类型名称为空,添加失败")
	}
	var adclass models.AdClass
	adclass.Class_name = classname
	adclass.Createdate = utils.GetTimeNow()
	b := models_ad_class.Save(adclass)
	if b == false {
		return ac.RenderJson("类型添加失败")
	}
	return ac.RenderJson("类型添加成功")
}

func (ac *AdClass) Delete() revel.Result {
	classid, err := strconv.Atoi(ac.Params.Get("classid"))
	if err != nil {
		return ac.RenderJson("类型删除失败")
	}
	models_ad_class := new(models.AdClass)
	b := models_ad_class.Delete(classid)
	if b == false {
		return ac.RenderJson("类型删除失败")
	}
	return ac.RenderJson("类型删除成功")
}
func (ac *AdClass) Update() revel.Result {
	classid, err := strconv.Atoi(ac.Params.Get("classid"))
	if err != nil {
		return ac.RenderJson("类型修改失败")
	}
	classname := ac.Params.Get("classname")
	models_ad_class := new(models.AdClass)
	var adclasses models.AdClass
	adclasses.Class_name = classname
	b := models_ad_class.Update(classid, adclasses)
	if b == false {
		return ac.RenderJson("类型修改失败")
	}
	return ac.RenderJson("类型修改成功")
}

func (ac *AdClass) Getall() revel.Result {
	model_ad_class := new(models.AdClass)
	var adclasses []models.AdClass
	adclasses = model_ad_class.Getall()
	return ac.RenderJson(adclasses)
}
