package controllers

import "github.com/revel/revel"
import "drinking_life/app/models"
import "drinking_life/app/utils"
import "strconv"

type Product struct {
	*revel.Controller
}

func (p *Product) Save() revel.Result {
	var product models.Product
	product.Productname = p.Params.Get("productname")
	price, _ := strconv.Atoi(p.Params.Get("price"))
	product.Price = price * 100
	product.Isframe = utils.Is_frame_down
	product.Sort_id, _ = strconv.Atoi(p.Params.Get("sortid"))
	product.Createtime = utils.GetTimeNow()

	pro := new(models.Product)
	b := pro.Save(product)
	if b == false {
		return p.RenderJson("产品添加失败")
	}
	return p.RenderJson("产品添加成功")
}

func (p *Product) Update() revel.Result {
	var product models.Product
	product.Productid, _ = strconv.Atoi(p.Params.Get("productid"))
	product.Productname = p.Params.Get("productname")
	price, _ := strconv.Atoi(p.Params.Get("price"))
	product.Price = price * 100
	product.Sort_id, _ = strconv.Atoi(p.Params.Get("sortid"))
	pro := new(models.Product)
	b := pro.Update(product)
	if b == false {
		return p.RenderJson("修改产品信息失败")
	}
	return p.RenderJson("修改产品信息成功")
}

func (p *Product) Upframe() revel.Result {
	productid, _ := strconv.Atoi(p.Params.Get("productid"))
	model_product := new(models.Product)
	b := model_product.Upframe(productid)
	if b == false {
		return p.RenderJson("产品上架失败")
	}
	return p.RenderJson("产品上架成功")
}

func (p *Product) Downfrmae() revel.Result {
	productid, _ := strconv.Atoi(p.Params.Get("productid"))
	model_product := new(models.Product)
	b := model_product.Downframe(productid)
	if b == false {
		return p.RenderJson("产品下架失败")
	}
	return p.RenderJson("产品上架成功")
}

func (p *Product) Getall() revel.Result {
	model_product := new(models.Product)
	var product []models.Product
	product = model_product.Getall()
	return p.RenderJson(product)
}
