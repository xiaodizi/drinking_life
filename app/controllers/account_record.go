package controllers

import "github.com/revel/revel"
import "drinking_life/app/models"
import "drinking_life/app/utils"
import "strconv"

type AccountRecord struct {
	*revel.Controller
}

func (ar *AccountRecord) Save() revel.Result {
	models_accountrecord := new(models.AccountRecord)
	reason, _ := strconv.Atoi(ar.Params.Get("reason"))
	amount, _ := strconv.Atoi(ar.Params.Get("amount"))
	ar_type, _ := strconv.Atoi(ar.Params.Get("type"))
	transfer_no := ar.Params.Get("transferno")
	account_id, _ := strconv.Atoi(ar.Params.Get("accountid"))
	if reason == 0 {
		return ar.RenderJson("变更原因为空,添加失败")
	}
	if amount == 0 {
		return ar.RenderJson("变更金额为空,添加失败")
	}
	if ar_type == 0 {
		return ar.RenderJson("账户类型为空,添加失败")
	}
	if transfer_no == "" {
		return ar.RenderJson("银行流水号为空,添加失败")
	}
	if account_id == 0 {
		return ar.RenderJson("账户id为空,添加失败")
	}
	var accountrecord models.AccountRecord
	accountrecord.Reason = reason
	accountrecord.Amount = amount
	accountrecord.Pubdate = utils.GetTimeNow()
	accountrecord.Type = ar_type
	accountrecord.Transfer_no = transfer_no
	accountrecord.Account_id = account_id
	b := models_accountrecord.Save(accountrecord)
	if b == false {
		return ar.RenderJson("账户变更添加失败")
	}
	return ar.RenderJson("账户变更添加成功")
}

func (ar *AccountRecord) Delete() revel.Result {
	recordid, err := strconv.Atoi(ar.Params.Get("recordid"))
	if err != nil {
		return ar.RenderJson("账户变更删除失败")
	}
	models_accountrecord := new(models.AccountRecord)
	b := models_accountrecord.Delete(recordid)
	if b == false {
		return ar.RenderJson("账户变更删除失败")
	}
	return ar.RenderJson("账户变更删除成功")
}

func (ar *AccountRecord) Update() revel.Result {
	recordid, _ := strconv.Atoi(ar.Params.Get("recordid"))
	reason, _ := strconv.Atoi(ar.Params.Get("reason"))
	amount, _ := strconv.Atoi(ar.Params.Get("amount"))
	ar_type, _ := strconv.Atoi(ar.Params.Get("type"))
	transfer_no := ar.Params.Get("transferno")
	account_id, _ := strconv.Atoi(ar.Params.Get("accountid"))
	models_accountrecord := new(models.AccountRecord)
	var accountrecord models.AccountRecord
	accountrecord.Reason = reason
	accountrecord.Amount = amount
	accountrecord.Pubdate = utils.GetTimeNow()
	accountrecord.Type = ar_type
	accountrecord.Transfer_no = transfer_no
	accountrecord.Account_id = account_id
	b := models_accountrecord.Update(recordid, accountrecord)
	if b == false {
		return ar.RenderJson("账户变更修改失败")
	}
	return ar.RenderJson("账户变更修改成功")
}

func (ar *AccountRecord) Getall() revel.Result {
	models_accountrecord := new(models.AccountRecord)
	var accountrecords []models.AccountRecord
	accountrecords = models_accountrecord.Getall()
	return ar.RenderJson(accountrecords)
}
