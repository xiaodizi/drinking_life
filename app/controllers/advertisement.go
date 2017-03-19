package controllers

import "github.com/revel/revel"
import "drinking_life/app/models"
import "drinking_life/app/utils"
import "strconv"

type Advertisement struct {
	*revel.Controller
}

func (ad *Advertisement) Save() revel.Result {
	models_advertisement := new(models.Advertisement)
	advertisementname := ad.Params.Get("advertisementname")
	photo := ad.Params.Get("photo")
	advertisementwords := ad.Params.Get("advertisementwords")
	isdisplay, _ := strconv.Atoi(ad.Params.Get("isdisplay"))
	sortid, _ := strconv.Atoi(ad.Params.Get("sortid"))
	classid, _ := strconv.Atoi(ad.Params.Get("classid"))
	if advertisementname == "" {
		return ad.RenderJson("广告名称为空,添加失败")
	}
	if photo == "" {
		return ad.RenderJson("广告图片为空,添加失败")
	}
	if advertisementwords == "" {
		return ad.RenderJson("广告语为空,添加失败")
	}
	if sortid == 0 {
		return ad.RenderJson("排序id为空,添加失败")
	}
	if classid == 0 {
		return ad.RenderJson("分类id为空,添加失败")
	}
	var advertisement models.Advertisement
	advertisement.Ad_name = advertisementname
	advertisement.Photo = photo
	advertisement.Ad_words = advertisementwords
	advertisement.Is_display = isdisplay
	advertisement.Sort_id = sortid
	advertisement.Class_id = classid
	advertisement.Createdate = utils.GetTimeNow()
	b := models_advertisement.Save(advertisement)
	if b == false {
		return ad.RenderJson("广告添加失败")
	}
	return ad.RenderJson("广告添加成功")
}

func (ad *Advertisement) Delete() revel.Result {
	adid, err := strconv.Atoi(ad.Params.Get("advertisementid"))
	if err != nil {
		return ad.RenderJson("广告删除失败")
	}
	models_ad := new(models.Advertisement)
	b := models_ad.Delete(adid)
	if b == false {
		return ad.RenderJson("广告删除失败")
	}
	return ad.RenderJson("广告删除成功")
}

func (ad *Advertisement) Update() revel.Result {
	advertisementid, _ := strconv.Atoi(ad.Params.Get("advertisementid"))
	advertisementname := ad.Params.Get("advertisementname")
	photo := ad.Params.Get("photo")
	advertisementwords := ad.Params.Get("advertisementwords")
	sortid, _ := strconv.Atoi(ad.Params.Get("sortid"))
	classid, _ := strconv.Atoi(ad.Params.Get("classid"))
	models_advertisement := new(models.Advertisement)
	var advertisements models.Advertisement
	advertisements.Ad_name = advertisementname
	advertisements.Photo = photo
	advertisements.Ad_words = advertisementwords
	advertisements.Sort_id = sortid
	advertisements.Class_id = classid
	b := models_advertisement.Update(advertisementid, advertisements)
	if b == false {
		return ad.RenderJson("广告修改失败")
	}
	return ad.RenderJson("广告修改成功")
}

func (ad *Advertisement) Getall() revel.Result {
	models_advertisement := new(models.Advertisement)
	var advertisements []models.Advertisement
	advertisements = models_advertisement.Getall()
	return ad.RenderJson(advertisements)
}
func (ad *Advertisement) Sort() revel.Result {
	count, _ := strconv.Atoi(ad.Params.Get("count"))
	models_advertisement := new(models.Advertisement)
	var advertisements []models.Advertisement
	advertisements = models_advertisement.Sort(count)
	return ad.RenderJson(advertisements)
}
func (ad *Advertisement) IsDisplay() revel.Result {
	ad_id, _ := strconv.Atoi(ad.Params.Get("ad_id"))
	model_advertisement := new(models.Advertisement)
	b := model_advertisement.IsDisplay(ad_id)
	if b == false {
		return ad.RenderJson("广告显示失败")
	}
	return ad.RenderJson("广告显示成功")
}
func (ad *Advertisement) IsnotDisplay() revel.Result {
	ad_id, _ := strconv.Atoi(ad.Params.Get("ad_id"))
	model_advertisement := new(models.Advertisement)
	b := model_advertisement.IsnotDisplay(ad_id)
	if b == false {
		return ad.RenderJson("广告取消显示失败")
	}
	return ad.RenderJson("广告取消显示成功")
}
